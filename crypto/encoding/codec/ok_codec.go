package codec

// ok codec used specifically for ok client

import (
	"github.com/ok-chain/ok-gosdk/types"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
)

type Codec = amino.Codec

var Cdc *Codec

func init() {
	cdc := New()
	// crypto register
	cryptoAmino.RegisterAmino(cdc)
	goSDKRegisterAmino(cdc)

	Cdc = cdc.Seal()

}

func New() *Codec {
	cdc := amino.NewCodec()
	return cdc
}

func goSDKRegisterAmino(cdc *amino.Codec){
	cdc.RegisterInterface((*types.Account)(nil), nil)
	cdc.RegisterConcrete(&types.BaseAccount{}, "auth/Account", nil)


}
