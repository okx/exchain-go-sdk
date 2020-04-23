package dex

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

var _ sdk.Module = (*dexClient)(nil)

type dexClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in dex module
func (dexClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (dexClient) Name() string {
	return types.ModuleName
}

// NewDexClient creates a new instance of dex client as implement
func NewDexClient(baseClient sdk.BaseClient) exposed.Dex {
	return dexClient{baseClient}
}
