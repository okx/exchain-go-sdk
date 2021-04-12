package token

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/token/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/x/token"
)

var _ gosdktypes.Module = (*tokenClient)(nil)

type tokenClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in token module
func (tokenClient) RegisterCodec(cdc *codec.Codec) {
	token.RegisterCodec(cdc)
}

// Name returns the module name
func (tokenClient) Name() string {
	return types.ModuleName
}

// NewTokenClient creates a new instance of token client as implement
func NewTokenClient(baseClient gosdktypes.BaseClient) exposed.Token {
	return tokenClient{baseClient}
}
