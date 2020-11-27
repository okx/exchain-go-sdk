package tendermint

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/tendermint/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

var _ gosdktypes.Module = (*tendermintClient)(nil)

type tendermintClient struct {
	gosdktypes.BaseClient
}

// nolint
func (tendermintClient) RegisterCodec(cdc gosdktypes.SDKCodec) {}

// Name returns the module name
func (tendermintClient) Name() string {
	return types.ModuleName
}

// NewTendermintClient creates a new instance of tendermint client as implement
func NewTendermintClient(baseClient gosdktypes.BaseClient) exposed.Tendermint {
	return tendermintClient{baseClient}
}
