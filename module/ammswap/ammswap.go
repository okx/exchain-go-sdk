package ammswap

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*ammswapClient)(nil)

type ammswapClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in ammswap module
func (ac ammswapClient) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (ammswapClient) Name() string {
	return types.ModuleName
}

// NewAmmSwapClient creates a new instance of ammswap client as implement
func NewAmmSwapClient(baseClient gosdktypes.BaseClient) exposed.AmmSwap {
	return ammswapClient{baseClient}
}
