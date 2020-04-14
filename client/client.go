package client

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/types"

	"github.com/okex/okchain-go-sdk/crypto/encoding/codec"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcCli "github.com/tendermint/tendermint/rpc/client"
)

var cdc *codec.Codec

func init() {
	cdc = codec.Cdc
}

type OKChainClient struct {
	rpcUrl string
	cli    *rpcCli.HTTP
	cdc    *codec.Codec
}

func NewClient(rpcUrl string) OKChainClient {
	return OKChainClient{
		rpcUrl: rpcUrl,
		cli:    rpcCli.NewHTTP(rpcUrl, "/websocket"),
		cdc:    cdc,
	}
}

func (cli *OKChainClient) querySubspace(subspace []byte, storeName string) (res []cmn.KVPair, err error) {
	resRaw, err := cli.queryStore(subspace, storeName, "subspace")
	if err != nil {
		return res, err
	}

	cli.cdc.MustUnmarshalBinaryLengthPrefixed(resRaw, &res)
	return
}

func (cli *OKChainClient) queryStore(key cmn.HexBytes, storeName, endPath string) ([]byte, error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, endPath)
	return cli.query(path, key)
}

func (cli *OKChainClient) query(path string, key cmn.HexBytes) ([]byte, error) {
	opts := rpcCli.ABCIQueryOptions{
		Height: 0,
		Prove:  false,
	}

	result, err := cli.cli.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return nil, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.New(resp.Log)
	}

	return resp.Value, nil

}

func (cli *OKChainClient) broadcast(txBytes []byte, broadcastMode string) (res types.TxResponse, err error) {
	switch broadcastMode {
	case BroadcastSync:
		res, err = doBroadcastTxSync(cli.cli, txBytes)

	case BroadcastAsync:
		res, err = doBroadcastTxAsync(cli.cli, txBytes)

	case BroadcastBlock:
		res, err = doBroadcastTxCommit(cli.cli, txBytes)
	default:
		err = fmt.Errorf("unsupported return broadcast mode %s; supported types: sync, async, block", broadcastMode)
	}
	return res, err
}

func doBroadcastTxSync(cli *rpcCli.HTTP, txBytes []byte) (types.TxResponse, error) {
	retBroadcastTx, err := cli.BroadcastTxSync(txBytes)
	return types.NewResponseFormatBroadcastTx(retBroadcastTx), err
}

func doBroadcastTxAsync(cli *rpcCli.HTTP, txBytes []byte) (types.TxResponse, error) {
	retBroadcastTx, err := cli.BroadcastTxAsync(txBytes)
	return types.NewResponseFormatBroadcastTx(retBroadcastTx), err
}

func doBroadcastTxCommit(cli *rpcCli.HTTP, txBytes []byte) (types.TxResponse, error) {
	retBroadcastTxCommit, err := cli.BroadcastTxCommit(txBytes)
	if err != nil {
		return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), err
	}
	if !retBroadcastTxCommit.CheckTx.IsOK() {
		return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), fmt.Errorf(retBroadcastTxCommit.CheckTx.Log)
	}
	if !retBroadcastTxCommit.DeliverTx.IsOK() {
		return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), fmt.Errorf(retBroadcastTxCommit.DeliverTx.Log)
	}
	return types.NewResponseFormatBroadcastTxCommit(retBroadcastTxCommit), nil
}
