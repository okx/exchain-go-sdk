package governance

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/governance/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/x/gov"
	paramstypes "github.com/okex/exchain/x/params/types"
)

var _ gosdktypes.Module = (*govClient)(nil)

type govClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in governance module
func (gc govClient) RegisterCodec(cdc *codec.Codec) {
	gov.RegisterCodec(cdc)
	cdc.RegisterConcrete(paramstypes.ParameterChangeProposal{}, "okexchain/params/ParameterChangeProposal", nil)
}

// Name returns the module name
func (gc govClient) Name() string {
	return types.ModuleName
}

// NewGovClient creates a new instance of governance client as implement
func NewGovClient(baseClient gosdktypes.BaseClient) exposed.Governance {
	return govClient{baseClient}
}
