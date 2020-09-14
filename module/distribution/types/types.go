package types

import (
	sdk "github.com/okex/okexchain-go-sdk/types"
)

// const
const (
	ModuleName = "distribution"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for distribution module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgSetWithdrawAddr{}, "okexchain/distribution/MsgModifyWithdrawAddress")
	cdc.RegisterConcrete(MsgWithdrawValCommission{}, "okexchain/distribution/MsgWithdrawReward")
}
