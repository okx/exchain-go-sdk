package types

import (
	cryptocodec "github.com/okx/okbchain/app/crypto/ethsecp256k1"
	evmtypes "github.com/okx/okbchain/app/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/codec"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
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
