package staking

import (
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/staking/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
)

var _ sdk.Module = (*stakingClient)(nil)

type stakingClient struct {
	sdk.BaseClient
}

// RegisterCodec registers the msg type in staking module
func (sc stakingClient) RegisterCodec(cdc sdk.SDKCodec) {
	types.RegisterCodec(cdc)
}

// Name returns the module name
func (stakingClient) Name() string {
	return types.ModuleName
}

// NewStakingClient creates a new instance of staking client as implement
func NewStakingClient(baseClient sdk.BaseClient) exposed.Staking {
	return stakingClient{baseClient}
}
