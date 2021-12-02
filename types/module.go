package types

import "github.com/okex/exchain/libs/cosmos-sdk/codec"

// Module shows the expected behaviour of each module in ExChain GoSDK
type Module interface {
	RegisterCodec(cdc *codec.Codec)
	Name() string
}
