package dex

import (
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/dex/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/x/dex"
)

var _ gosdktypes.Module = (*dexClient)(nil)

type dexClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in dex module
func (dexClient) RegisterCodec(cdc *codec.Codec) {
	dex.RegisterCodec(cdc)
}

// Name returns the module name
func (dexClient) Name() string {
	return types.ModuleName
}

// NewDexClient creates a new instance of dex client as implement
func NewDexClient(baseClient gosdktypes.BaseClient) exposed.Dex {
	return dexClient{baseClient}
}
