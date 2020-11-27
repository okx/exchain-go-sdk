package ammswap

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Module = (*ammswapClient)(nil)

type ammswapClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in ammswap module
func (pc ammswapClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (ammswapClient) Name() string {
	return types.ModuleName
}

// NewAmmSwapClient creates a new instance of ammswap client as implement
func NewAmmSwapClient(baseClient sdk.BaseClient) exposed.AmmSwap {
	return ammswapClient{baseClient}
}
