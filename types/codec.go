package types

import (
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/multisig"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

// Cdc is the codec for signing
var Cdc = NewCodec()

// SDKCodec shows the expected behaviour of codec in okchain gosdk
type SDKCodec interface {
	MarshalJSON(o interface{}) ([]byte, error)
	UnmarshalJSON(bytes []byte, ptr interface{}) error
	MustMarshalJSON(o interface{}) []byte
	MustUnmarshalJSON(bytes []byte, ptr interface{})

	MarshalBinaryLengthPrefixed(o interface{}) ([]byte, error)
	UnmarshalBinaryLengthPrefixed(bytes []byte, ptr interface{}) error
	MustUnmarshalBinaryLengthPrefixed(bytes []byte, ptr interface{})

	UnmarshalBinaryBare(bytes []byte, ptr interface{}) error

	RegisterConcrete(o interface{}, name string)
	RegisterInterface(ptr interface{})

	Seal()
}

var _ SDKCodec = (*Codec)(nil)

// SDKCodec defines the codec only for okchain gosdk
type Codec struct {
	*amino.Codec
}

// NewCodec creates a new instance of codec only for gosdk
func NewCodec() SDKCodec {
	return Codec{amino.NewCodec()}
}

// RegisterConcrete implements the SDKCodec interface
func (cdc Codec) RegisterConcrete(o interface{}, name string) {
	cdc.Codec.RegisterConcrete(o, name, nil)
}

// RegisterInterface implements the SDKCodec interface
func (cdc Codec) RegisterInterface(ptr interface{}) {
	cdc.Codec.RegisterInterface(ptr, nil)
}

// Seal implements the SDKCodec interface
func (cdc Codec) Seal() {
	cdc.Codec.Seal()
}

// RegisterBasicCodec registers the basic data types for gosdk codec
func RegisterBasicCodec(cdc SDKCodec) {
	// amino
	cdc.RegisterInterface((*crypto.PubKey)(nil))
	cdc.RegisterConcrete(ed25519.PubKeyEd25519{}, ed25519.PubKeyAminoName)
	cdc.RegisterConcrete(secp256k1.PubKeySecp256k1{}, secp256k1.PubKeyAminoName)
	cdc.RegisterConcrete(multisig.PubKeyMultisigThreshold{}, multisig.PubKeyMultisigThresholdAminoRoute)
	cdc.RegisterInterface((*crypto.PrivKey)(nil))
	cdc.RegisterConcrete(ed25519.PrivKeyEd25519{}, ed25519.PrivKeyAminoName)
	cdc.RegisterConcrete(secp256k1.PrivKeySecp256k1{}, secp256k1.PrivKeyAminoName)
	// stdTx
	cdc.RegisterInterface((*Tx)(nil))
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx")
	// msg
	cdc.RegisterInterface((*Msg)(nil))
}
