package tendermint

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/tendermint/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
)

var _ gosdktypes.Module = (*tendermintClient)(nil)

type tendermintClient struct {
	gosdktypes.BaseClient
}

// nolint
func (tendermintClient) RegisterCodec(_ *codec.Codec) {}

// Name returns the module name
func (tendermintClient) Name() string {
	return types.ModuleName
}

// NewTendermintClient creates a new instance of tendermint client as implement
func NewTendermintClient(baseClient gosdktypes.BaseClient) exposed.Tendermint {
	return tendermintClient{baseClient}
}
