package backend

import (
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/backend/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
)

var _ gosdktypes.Module = (*backendClient)(nil)

type backendClient struct {
	gosdktypes.BaseClient
}

// nolint
func (backendClient) RegisterCodec(_ *codec.Codec) {}

// Name returns the module name
func (backendClient) Name() string {
	return types.ModuleName
}

// NewBackendClient creates a new instance of backend client as implement
func NewBackendClient(baseClient gosdktypes.BaseClient) exposed.Backend {
	return backendClient{baseClient}
}
