package okclient

import (
	"github.com/ok-chain/ok-gosdk/common/libs/cosmos/cosmos-sdk/codec"
	rpcCli "github.com/tendermint/tendermint/rpc/client"
)

var (
	cdc = codec.New()
)

func init() {

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
