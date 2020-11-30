package dex

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/dex/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*dexClient)(nil)

type dexClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in dex module
func (dexClient) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (dexClient) Name() string {
	return types.ModuleName
}

// NewDexClient creates a new instance of dex client as implement
func NewDexClient(baseClient gosdktypes.BaseClient) exposed.Dex {
	return dexClient{baseClient}
}
