package types

import (
	sdk "github.com/okex/okexchain-go-sdk/types"
)

// const
const (
	ModuleName = "ammswap"
)

var (
	msgCdc = sdk.NewCodec()
	// TokenPairPrefixKey to be used for KVStore
	TokenPairPrefixKey = []byte{0x01}
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for ammswap module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgAddLiquidity{}, "okexchain/ammswap/MsgAddLiquidity")
	cdc.RegisterConcrete(MsgRemoveLiquidity{}, "okexchain/ammswap/MsgRemoveLiquidity")
	cdc.RegisterConcrete(MsgCreateExchange{}, "okexchain/ammswap/MsgCreateExchange")
	cdc.RegisterConcrete(MsgTokenToNativeToken{}, "okexchain/ammswap/MsgSwapToken")
}

// SwapTokenPair defines token pair exchange
type SwapTokenPair struct {
	QuotePooledCoin sdk.DecCoin `json:"quote_pooled_coin"` // The volume of quote token in the token pair exchange pool
	BasePooledCoin  sdk.DecCoin `json:"base_pooled_coin"`  // The volume of base token in the token pair exchange pool
	PoolTokenName   string      `json:"pool_token_name"`   // The name of pool token
}

// nolint
func GetTokenPairKey(key string) []byte {
	return append(TokenPairPrefixKey, []byte(key)...)
}