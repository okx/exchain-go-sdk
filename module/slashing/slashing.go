package slashing

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/module/slashing/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

var _ sdk.Module = (*slashingClient)(nil)

type slashingClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in slashing module
func (sc slashingClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (slashingClient) Name() string {
	return types.ModuleName
}

// NewSlashingClient creates a new instance of slashing client as implement
func NewSlashingClient(baseClient sdk.BaseClient) exposed.Slashing {
	return slashingClient{baseClient}
}
