package token

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module/token/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

var _ sdk.Module = (*tokenClient)(nil)

type tokenClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in token module
func (tokenClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (tokenClient) Name() string {
	return types.ModuleName
}

// NewTokenClient creates a new instance of token client as implement
func NewTokenClient(baseClient sdk.BaseClient) exposed.Token {
	return tokenClient{baseClient}
}
