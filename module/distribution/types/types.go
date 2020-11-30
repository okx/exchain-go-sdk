package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/distribution"
)

// const
const (
	ModuleName = "distribution"
)

var (
	msgCdc = gosdktypes.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for distribution module
func RegisterCodec(cdc *codec.Codec) {
	distribution.RegisterCodec(cdc)
}
