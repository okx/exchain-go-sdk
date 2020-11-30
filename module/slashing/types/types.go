package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/slashing"
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
func RegisterCodec(cdc *codec.Codec) {
	slashing.RegisterCodec(cdc)
}
