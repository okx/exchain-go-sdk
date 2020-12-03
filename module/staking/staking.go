package staking

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/okex/okexchain-go-sdk/exposed"
	"github.com/okex/okexchain-go-sdk/module/staking/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/staking"
)

var _ gosdktypes.Module = (*stakingClient)(nil)

type stakingClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in staking module
func (sc stakingClient) RegisterCodec(cdc *codec.Codec) {
	staking.RegisterCodec(cdc)
}

// Name returns the module name
func (stakingClient) Name() string {
	return types.ModuleName
}

// NewStakingClient creates a new instance of staking client as implement
func NewStakingClient(baseClient gosdktypes.BaseClient) exposed.Staking {
	return stakingClient{baseClient}
}
