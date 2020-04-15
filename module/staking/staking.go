package staking

import (
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
)

var _ types.Module = (*stakingClient)(nil)

type stakingClient struct {
	types.BaseClient
}

// RegisterCodec registers the msg type in staking module
func (sc stakingClient) RegisterCodec(cdc types.SDKCodec) {
	registerCodec(cdc)
}

// Name returns the module name
func (stakingClient) Name() string {
	return ModuleName
}

// NewStakingClient creates a new instance of staking client as implement
func NewStakingClient(baseClient types.BaseClient) exposed.Staking {
	return stakingClient{baseClient}
}

func registerCodec(cdc types.SDKCodec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "okchain/staking/MsgCreateValidator")
	cdc.RegisterConcrete(MsgEditValidator{}, "okchain/staking/MsgEditValidator")
	cdc.RegisterConcrete(MsgDelegate{}, "okchain/staking/MsgDelegate")
	cdc.RegisterConcrete(MsgUndelegate{}, "okchain/staking/MsgUnDelegate")
	cdc.RegisterConcrete(MsgVote{}, "okchain/staking/MsgVote")
	cdc.RegisterConcrete(MsgDestroyValidator{}, "okchain/staking/MsgDestroyValidator")
	cdc.RegisterConcrete(MsgRegProxy{}, "okchain/staking/MsgRegProxy")
	cdc.RegisterConcrete(MsgBindProxy{}, "okchain/staking/MsgBindProxy")
	cdc.RegisterConcrete(MsgUnbindProxy{}, "okchain/staking/MsgUnbindProxy")
}
