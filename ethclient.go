package gosdk

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type rpcClient struct {
	*rpc.Client
}

type ethClient struct {
	*ethclient.Client
	*rpcClient
}
func NewEthClient(ctx context.Context, rawurl string) (*ethClient, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return &ethClient{ethclient.NewClient(c), &rpcClient{c}}, nil
}
