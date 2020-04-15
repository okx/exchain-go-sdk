package slashing

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
)

var _ types.Module = (*slashingClient)(nil)

type slashingClient struct {
	types.BaseClient
}

// RegisterCodec registers the msg type in slashing module
func (sc slashingClient) RegisterCodec(cdc types.SDKCodec) {
	registerCodec(cdc)
}

// Name returns the module name
func (slashingClient) Name() string {
	return ModuleName
}

// NewSlashingClient creates a new instance of slashing client as implement
func NewSlashingClient(baseClient types.BaseClient) exposed.Slashing {
	return slashingClient{baseClient}
}

func registerCodec(cdc types.SDKCodec) {
	cdc.RegisterConcrete(MsgUnjail{}, "cosmos-sdk/MsgUnjail")
}
