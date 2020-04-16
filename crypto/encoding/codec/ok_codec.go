package codec

// ok codec used specifically for ok client

import (
	"github.com/okex/okchain-go-sdk/types"
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
	return amino.NewCodec()
}

func goSDKRegisterAmino(cdc *amino.Codec) {
	//cdc.RegisterInterface((*types.Account)(nil), nil)
	//cdc.RegisterConcrete(&types.BaseAccount{}, "cosmos-sdk/Account", nil)
	cdc.RegisterInterface((*types.Proposal)(nil), nil)
	cdc.RegisterConcrete(&types.TextProposal{}, "okchain/gov/TextProposal", nil)
	cdc.RegisterConcrete(&types.DexListProposal{}, "okchain/gov/DexListProposal", nil)
	cdc.RegisterConcrete(&types.ParameterProposal{}, "okchain/gov/ParameterProposal", nil)
	cdc.RegisterConcrete(&types.AppUpgradeProposal{}, "okchain/gov/AppUpgradeProposal", nil)

}



