package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/ammswap"
)

// const
const (
	ModuleName = ammswap.ModuleName
)

var (
	msgCdc = gosdktypes.NewCodec()
	// TokenPairPrefixKey to be used for KVStore
	TokenPairPrefixKey = []byte{0x01}
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for ammswap module
func RegisterCodec(cdc *codec.Codec) {
	ammswap.RegisterCodec(cdc)
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

// nolint
func GetSwapTokenPairName(token1, token2 string) string {
	if token1 < token2 {
		return token1 + "_" + token2
	}
	return token2 + "_" + token1
}
