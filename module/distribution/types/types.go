package types

import (
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// const
const (
	ModuleName = "distribution"
)

var (
	msgCdc = gosdktypes.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for distribution module
func RegisterCodec(cdc gosdktypes.SDKCodec) {
	cdc.RegisterConcrete(MsgSetWithdrawAddr{}, "okexchain/distribution/MsgModifyWithdrawAddress")
	cdc.RegisterConcrete(MsgWithdrawValCommission{}, "okexchain/distribution/MsgWithdrawReward")
}
