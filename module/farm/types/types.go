package types

import sdk "github.com/okex/okexchain-go-sdk/types"

// const
const (
	ModuleName = "farm"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for farm module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "okexchain/farm/MsgCreatePool")
	cdc.RegisterConcrete(MsgDestroyPool{}, "okexchain/farm/MsgDestroyPool")
	//cdc.RegisterConcrete(MsgLock{}, "okexchain/farm/MsgLock")
	//cdc.RegisterConcrete(MsgUnlock{}, "okexchain/farm/MsgUnlock")
	//cdc.RegisterConcrete(MsgClaim{}, "okexchain/farm/MsgClaim")
	cdc.RegisterConcrete(MsgProvide{}, "okexchain/farm/MsgProvide")
}
