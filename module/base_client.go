package module

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcCli "github.com/tendermint/tendermint/rpc/client"
)

var _ types.BaseClient = (*baseClient)(nil)

type baseClient struct {
	types.RPCClient
	config *types.ClientConfig
	cdc    types.SDKCodec
}

// NewBaseClient creates a new instance of baseClient
func NewBaseClient(cdc types.SDKCodec, pConfig *types.ClientConfig) *baseClient {
	return &baseClient{
		RPCClient: rpcCli.NewHTTP(pConfig.NodeURI, "/websocket"),
		config:    pConfig,
		cdc:       cdc,
	}
}

// Query executes the basic query
func (bc *baseClient) Query(path string, key cmn.HexBytes) ([]byte, error) {
	opts := rpcCli.ABCIQueryOptions{
		Height: 0,
		Prove:  false,
	}

	result, err := bc.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return nil, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.New(resp.Log)
	}

	return resp.Value, nil
}

// QueryStore executes the direct query to the store
func (bc *baseClient) QueryStore(key cmn.HexBytes, storeName, endPath string) ([]byte, error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, endPath)
	return bc.Query(path, key)
}

// QuerySubspace executes the direct query to the subspace
func (bc *baseClient) QuerySubspace(subspace []byte, storeName string) (res []cmn.KVPair, err error) {
	resRaw, err := bc.QueryStore(subspace, storeName, "subspace")
	if err != nil {
		return
	}

	bc.cdc.MustUnmarshalBinaryLengthPrefixed(resRaw, &res)
	return
}

// Broadcast broadcasts by different modes
func (bc *baseClient) Broadcast(txBytes []byte, broadcastMode types.BroadcastMode) (res types.TxResponse, err error) {
	switch broadcastMode {
	case types.BroadcastSync:
		retBroadcastTx, err := bc.BroadcastTxSync(txBytes)
		return types.NewResponseFormatBroadcastTx(retBroadcastTx), err

	case types.BroadcastAsync:
		retBroadcastTx, err := bc.BroadcastTxAsync(txBytes)
		return types.NewResponseFormatBroadcastTx(retBroadcastTx), err

	case types.BroadcastBlock:
		retBroadcastTxCommit, err := bc.BroadcastTxCommit(txBytes)
		if err != nil {
			return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), err
		}
		if !retBroadcastTxCommit.CheckTx.IsOK() {
			return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), fmt.Errorf(retBroadcastTxCommit.CheckTx.Log)
		}
		if !retBroadcastTxCommit.DeliverTx.IsOK() {
			return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), fmt.Errorf(retBroadcastTxCommit.DeliverTx.Log)
		}
		return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), err

	default:
		err = fmt.Errorf("failed. unsupported broadcast mode %s; supported types: sync, async, block", broadcastMode)
	}
	return
}

// GetCodec gets the codec of the base client
func (bc *baseClient) GetCodec() types.SDKCodec {
	return bc.cdc
}

// GetConfig gets the client config
func (bc *baseClient) GetConfig() types.ClientConfig {
	return *bc.config
}

// BuildAndBroadcast implements the TxHandler interface
func (bc *baseClient) BuildAndBroadcast(fromName, passphrase, memo string, msgs []types.Msg, accNumber,
	seqNumber uint64) (resp types.TxResponse, err error) {
	stdTx, err := tx.BuildTx(fromName, passphrase, memo, msgs, accNumber, seqNumber)
	if err != nil {
		return resp, fmt.Errorf("failed. build stdTx error: %s", err)
	}

	bytes, err := bc.cdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return resp, fmt.Errorf("failed. encoded stdTx error: %s", err)
	}

	return bc.Broadcast(bytes, bc.GetConfig().BroadcastMode)
}
