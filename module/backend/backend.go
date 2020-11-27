package backend

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/backend/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Module = (*backendClient)(nil)

type backendClient struct {
	sdk.BaseClient
}

// nolint
func (backendClient) RegisterCodec(cdc sdk.SDKCodec) {}

// Name returns the module name
func (backendClient) Name() string {
	return types.ModuleName
}

// NewBackendClient creates a new instance of backend client as implement
func NewBackendClient(baseClient sdk.BaseClient) exposed.Backend {
	return backendClient{baseClient}
}
