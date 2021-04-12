package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cryptocodec "github.com/okex/exchain/app/crypto/ethsecp256k1"
	evmtypes "github.com/okex/exchain/app/types"
)

// NewCodec creates a new instance of codec only for gosdk
func NewCodec() *codec.Codec {
	return codec.New()
}

// RegisterBasicCodec registers the basic data types for gosdk codec
func RegisterBasicCodec(cdc *codec.Codec) {
	sdk.RegisterCodec(cdc)
	cryptocodec.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	evmtypes.RegisterCodec(cdc)
}
