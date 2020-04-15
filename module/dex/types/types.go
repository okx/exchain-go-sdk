package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "dex"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for dex module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgList{}, "okchain/dex/MsgList")
	//cdc.RegisterConcrete(MsgDeposit{}, "okchain/dex/MsgDeposit")
	//cdc.RegisterConcrete(MsgWithdraw{}, "okchain/dex/MsgWithdraw")
	//cdc.RegisterConcrete(MsgTransferOwnership{}, "okchain/dex/MsgTransferTradingPairOwnership")
}
