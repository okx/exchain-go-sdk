package types

import sdk "github.com/okex/okexchain-go-sdk/types"

// const
const (
	ModuleName = "slashing"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for slashing module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgUnjail{}, "cosmos-sdk/MsgUnjail")
}
