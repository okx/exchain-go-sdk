package ammswap

import (
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/ammswap/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/x/ammswap"
)

var _ gosdktypes.Module = (*ammswapClient)(nil)

type ammswapClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in ammswap module
func (ac ammswapClient) RegisterCodec(cdc *codec.Codec) {
	ammswap.RegisterCodec(cdc)
}

// Name returns the module name
func (ammswapClient) Name() string {
	return types.ModuleName
}

// NewAmmSwapClient creates a new instance of ammswap client as implement
func NewAmmSwapClient(baseClient gosdktypes.BaseClient) exposed.AmmSwap {
	return ammswapClient{baseClient}
}
