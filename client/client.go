package client

import (
	"errors"
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
