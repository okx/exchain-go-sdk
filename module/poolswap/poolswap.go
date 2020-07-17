package poolswap

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module/poolswap/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

var _ sdk.Module = (*poolswapClient)(nil)

type poolswapClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in staking module
func (pc poolswapClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (poolswapClient) Name() string {
	return types.ModuleName
}

// NewStakingClient creates a new instance of staking client as implement
func NewPoolSwapClient(baseClient sdk.BaseClient) exposed.PoolSwap {
	return poolswapClient{baseClient}
}