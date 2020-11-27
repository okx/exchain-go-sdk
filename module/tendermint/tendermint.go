package tendermint

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Module = (*tendermintClient)(nil)

type tendermintClient struct {
	sdk.BaseClient
}

// nolint
func (tendermintClient) RegisterCodec(cdc sdk.SDKCodec) {}

// Name returns the module name
func (tendermintClient) Name() string {
	return types.ModuleName
}

// NewTendermintClient creates a new instance of tendermint client as implement
func NewTendermintClient(baseClient sdk.BaseClient) exposed.Tendermint {
	return tendermintClient{baseClient}
}
