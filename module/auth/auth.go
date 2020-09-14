package auth

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/auth/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
)

var _ sdk.Module = (*authClient)(nil)

type authClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the account type in auth module
func (authClient) RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterInterface((*types.Account)(nil))
	cdc.RegisterConcrete(&types.BaseAccount{}, "cosmos-sdk/Account")
}

// Name returns the module name
func (authClient) Name() string {
	return types.ModuleName
}

// NewAuthClient creates a new instance of auth client as implement
func NewAuthClient(baseClient sdk.BaseClient) exposed.Auth {
	return authClient{baseClient}
}
