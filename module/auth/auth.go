package auth

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
)

var _ types.Module = (*authClient)(nil)

type authClient struct {
	types.BaseClient
}

// RegisterCodec registers the account type in auth module
func (authClient) RegisterCodec(cdc types.SDKCodec) {
	cdc.RegisterInterface((*exposed.Account)(nil))
	cdc.RegisterConcrete(&exposed.BaseAccount{}, "cosmos-sdk/Account")
}

// Name returns the module name
func (authClient) Name() string {
	return ModuleName
}

func NewAuthClient(baseClient types.BaseClient) exposed.Auth {
	return authClient{baseClient}
}
