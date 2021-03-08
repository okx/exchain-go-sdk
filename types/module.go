package types

import "github.com/cosmos/cosmos-sdk/codec"

// Module shows the expected behaviour of each module in OKExChain GoSDK
type Module interface {
	RegisterCodec(cdc *codec.Codec)
	Name() string
}
