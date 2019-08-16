package okclient

import (
	"errors"
	"github.com/ok-chain/ok-gosdk/crypto/encoding/codec"
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
