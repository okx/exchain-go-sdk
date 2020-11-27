package auth

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/auth/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*authClient)(nil)

type authClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the account type in auth module
func (authClient) RegisterCodec(cdc gosdktypes.SDKCodec) {
	cdc.RegisterInterface((*types.Account)(nil))
	cdc.RegisterConcrete(&authtypes.BaseAccount{}, "cosmos-sdk/Account")
}

// Name returns the module name
func (authClient) Name() string {
	return types.ModuleName
}

// NewAuthClient creates a new instance of auth client as implement
func NewAuthClient(baseClient gosdktypes.BaseClient) exposed.Auth {
	return authClient{baseClient}
}
