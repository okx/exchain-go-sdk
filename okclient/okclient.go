package okclient

import (
	"errors"
	"fmt"
	"github.com/ok-chain/gosdk/types"

	"github.com/ok-chain/gosdk/crypto/encoding/codec"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcCli "github.com/tendermint/tendermint/rpc/client"
)

var (
	cdc *codec.Codec
)

func init() {
	cdc = codec.Cdc
}

type OKClient struct {
	rpcUrl string
	cli    *rpcCli.HTTP
	cdc    *codec.Codec
}

func NewClient(rpcUrl string) OKClient {
	return OKClient{
		rpcUrl: rpcUrl,
		cli:    rpcCli.NewHTTP(rpcUrl, "/websocket"),
		cdc:    cdc,
	}
}

func (okCli *OKClient) query(path string, key cmn.HexBytes) ([]byte, error) {
	opts := rpcCli.ABCIQueryOptions{
		Height: 0,
		Prove:  false,
	}

	result, err := okCli.cli.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return nil, err
	}

	resp := result.Response
	if !resp.IsOK() {
		return nil, errors.New(resp.Log)
	}

	return resp.Value, nil

}

func (okCli *OKClient) broadcast(txBytes []byte, broadcastMode string) (res types.TxResponse, err error) {
	switch broadcastMode {
	case BroadcastSync:
		res, err = doBroadcastTxSync(okCli.cli, txBytes)

	case BroadcastAsync:
		res, err = doBroadcastTxAsync(okCli.cli, txBytes)

	case BroadcastBlock:
		res, err = doBroadcastTxCommit(okCli.cli, txBytes)
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
