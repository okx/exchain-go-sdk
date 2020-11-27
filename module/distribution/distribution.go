package distribution

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/distribution/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*distrClient)(nil)

type distrClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in distribution module
func (dc distrClient) RegisterCodec(cdc gosdktypes.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (distrClient) Name() string {
	return types.ModuleName
}

// NewDistrClient creates a new instance of distribution client as implement
func NewDistrClient(baseClient gosdktypes.BaseClient) exposed.Distribution {
	return distrClient{baseClient}
}
