package types

import (
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// const
const (
	ModuleName = "slashing"
)

var (
	msgCdc = gosdktypes.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for slashing module
func RegisterCodec(cdc gosdktypes.SDKCodec) {
	cdc.RegisterConcrete(MsgUnjail{}, "cosmos-sdk/MsgUnjail")
}
