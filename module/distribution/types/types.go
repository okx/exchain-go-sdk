package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
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
	cdc.RegisterConcrete(MsgSetWithdrawAddr{}, "okchain/distribution/MsgModifyWithdrawAddress")
}
