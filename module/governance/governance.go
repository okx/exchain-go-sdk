package governance

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/governance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Module = (*govClient)(nil)

type govClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in governance module
func (gc govClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (gc govClient) Name() string {
	return types.ModuleName
}

// NewGovClient creates a new instance of governance client as implement
func NewGovClient(baseClient sdk.BaseClient) exposed.Governance {
	return govClient{baseClient}
}
