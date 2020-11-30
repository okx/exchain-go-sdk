package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cryptocodec "github.com/okex/okexchain/app/crypto/ethsecp256k1"
	evmtypes "github.com/okex/okexchain/app/types"
	"github.com/tendermint/go-amino"
)

// Cdc is the codec for signing
var Cdc = NewCodec()

//// SDKCodec shows the expected behaviour of codec in okexchain gosdk
type SDKCodec interface {
	MarshalJSON(o interface{}) ([]byte, error)
	UnmarshalJSON(bytes []byte, ptr interface{}) error
	MustMarshalJSON(o interface{}) []byte
	MustUnmarshalJSON(bytes []byte, ptr interface{})

	MarshalBinaryLengthPrefixed(o interface{}) ([]byte, error)
	UnmarshalBinaryLengthPrefixed(bytes []byte, ptr interface{}) error
	MustMarshalBinaryLengthPrefixed(o interface{}) []byte
	MustUnmarshalBinaryLengthPrefixed(bytes []byte, ptr interface{})

	MarshalBinaryBare(o interface{}) ([]byte, error)
	UnmarshalBinaryBare(bytes []byte, ptr interface{}) error

	RegisterConcrete(o interface{}, name string, copts *amino.ConcreteOptions)
	RegisterInterface(ptr interface{}, iopts *amino.InterfaceOptions)

	Seal() *amino.Codec
}

//var _ SDKCodec = (*Codec)(nil)

//// Codec defines the codec only for OKExChain gosdk
//type Codec struct {
//	*amino.Codec
//}

// NewCodec creates a new instance of codec only for gosdk
func NewCodec() *codec.Codec {
	return codec.New()
}
//
//// RegisterConcrete implements the SDKCodec interface
//func (cdc Codec) RegisterConcrete(o interface{}, name string, copts *amino.ConcreteOptions) {
//	cdc.Codec.RegisterConcrete(o, name, copts)
//}
//
//// RegisterInterface implements the SDKCodec interface
//func (cdc Codec) RegisterInterface(ptr interface{}, iopts *amino.InterfaceOptions) {
//	cdc.Codec.RegisterInterface(ptr, iopts)
//}
//
//// Seal implements the SDKCodec interface
//func (cdc Codec) Seal() *amino.Codec {
//	return cdc.Codec.Seal()
//}

// RegisterBasicCodec registers the basic data types for gosdk codec
func RegisterBasicCodec(cdc *codec.Codec) {
	sdk.RegisterCodec(cdc)
	cryptocodec.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	evmtypes.RegisterCodec(cdc)
}
