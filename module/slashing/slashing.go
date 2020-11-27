package slashing

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/slashing/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*slashingClient)(nil)

type slashingClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in slashing module
func (sc slashingClient) RegisterCodec(cdc gosdktypes.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (slashingClient) Name() string {
	return types.ModuleName
}

// NewSlashingClient creates a new instance of slashing client as implement
func NewSlashingClient(baseClient gosdktypes.BaseClient) exposed.Slashing {
	return slashingClient{baseClient}
}
