package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "poolswap"
)

var (
	msgCdc = sdk.NewCodec()
	// TokenPairPrefixKey to be used for KVStore
	TokenPairPrefixKey = []byte{0x01}
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for staking module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgAddLiquidity{}, "okchain/poolswap/MsgAddLiquidity")
	cdc.RegisterConcrete(MsgRemoveLiquidity{}, "okchain/poolswap/MsgRemoveLiquidity")
	cdc.RegisterConcrete(MsgCreateExchange{}, "okchain/poolswap/MsgCreateExchange")
	cdc.RegisterConcrete(MsgTokenToNativeToken{}, "okchain/poolswap/MsgSwapToken")
}