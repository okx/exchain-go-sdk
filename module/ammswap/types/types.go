package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/ammswap"
	ammswaptypes "github.com/okex/okexchain/x/ammswap/types"
)

// const
const (
	ModuleName = ammswap.ModuleName
)

var (
	msgCdc = gosdktypes.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for ammswap module
func RegisterCodec(cdc *codec.Codec) {
	ammswap.RegisterCodec(cdc)
}

type (
	SwapTokenPair = ammswaptypes.SwapTokenPair
)

// nolint
func GetSwapTokenPairName(token1, token2 string) string {
	if token1 < token2 {
		return token1 + "_" + token2
	}
	return token2 + "_" + token1
}
