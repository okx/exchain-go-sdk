package module

import (
	"errors"
	"fmt"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/tx"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcCli "github.com/tendermint/tendermint/rpc/client"
)

var _ sdk.BaseClient = (*baseClient)(nil)

type baseClient struct {
	sdk.RPCClient
	config *sdk.ClientConfig
	cdc    sdk.SDKCodec
}

// NewBaseClient creates a new instance of baseClient
func NewBaseClient(cdc sdk.SDKCodec, pConfig *sdk.ClientConfig) *baseClient {
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
func (bc *baseClient) Broadcast(txBytes []byte, broadcastMode sdk.BroadcastMode) (res sdk.TxResponse, err error) {
	switch broadcastMode {
	case sdk.BroadcastSync:
		retBroadcastTx, err := bc.BroadcastTxSync(txBytes)
		return sdk.NewResponseFormatBroadcastTx(retBroadcastTx), err

	case sdk.BroadcastAsync:
		retBroadcastTx, err := bc.BroadcastTxAsync(txBytes)
		return sdk.NewResponseFormatBroadcastTx(retBroadcastTx), err

	case sdk.BroadcastBlock:
		retBroadcastTxCommit, err := bc.BroadcastTxCommit(txBytes)
		if err != nil {
			return sdk.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), err
		}
		if !retBroadcastTxCommit.CheckTx.IsOK() {
			return sdk.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), fmt.Errorf(retBroadcastTxCommit.CheckTx.Log)
		}
		if !retBroadcastTxCommit.DeliverTx.IsOK() {
			return sdk.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), fmt.Errorf(retBroadcastTxCommit.DeliverTx.Log)
		}
		return sdk.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), err

	default:
		err = fmt.Errorf("failed. unsupported broadcast mode %s; supported types: sync, async, block", broadcastMode)
	}
	return
}

// GetCodec gets the codec of the base client
func (bc *baseClient) GetCodec() sdk.SDKCodec {
	return bc.cdc
}

// GetConfig gets the client config
func (bc *baseClient) GetConfig() sdk.ClientConfig {
	return *bc.config
}

// BuildAndBroadcast implements the TxHandler interface
func (bc *baseClient) BuildAndBroadcast(fromName, passphrase, memo string, msgs []sdk.Msg, accNumber,
	seqNumber uint64) (resp sdk.TxResponse, err error) {
	stdTx, err := bc.BuildStdTx(fromName, passphrase, memo, msgs, accNumber, seqNumber)
	if err != nil {
		return resp, fmt.Errorf("failed. build stdTx error: %s", err)
	}

	bytes, err := bc.cdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return resp, fmt.Errorf("failed. encoded stdTx error: %s", err)
	}

	return bc.Broadcast(bytes, bc.GetConfig().BroadcastMode)
}

// BuildAndSign builds std sign context and sign it
func (bc *baseClient) BuildStdTx(fromName, passphrase, memo string, msgs []sdk.Msg, accNumber, seqNumber uint64) (
	stdTx sdk.StdTx, err error) {
	config := bc.GetConfig()
	if len(config.ChainID) == 0 {
		return stdTx, errors.New("failed. empty chain ID")
	}
	// TODO: implements the gas price later
	signMsg := sdk.StdSignMsg{
		ChainID:       config.ChainID,
		AccountNumber: accNumber,
		Sequence:      seqNumber,
		Memo:          memo,
		Msgs:          msgs,
		Fee:           sdk.NewStdFee(200000, config.Fees),
	}

	sigBytes, err := tx.MakeSignature(fromName, passphrase, signMsg)
	if err != nil {
		return
	}

	return sdk.NewStdTx(signMsg.Msgs, signMsg.Fee, []sdk.StdSignature{sigBytes}, signMsg.Memo), err
}
