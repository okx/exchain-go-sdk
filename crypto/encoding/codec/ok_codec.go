package codec

// ok codec used specifically for ok client

import (
	"encoding/json"
	"github.com/ok-chain/gosdk/common"
	"github.com/ok-chain/gosdk/types"
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

func goSDKRegisterAmino(cdc *amino.Codec) {
	cdc.RegisterInterface((*types.Account)(nil), nil)
	cdc.RegisterConcrete(&types.BaseAccount{}, "auth/Account", nil)
	cdc.RegisterInterface((*types.Proposal)(nil), nil)
	cdc.RegisterConcrete(&types.TextProposal{}, "okchain/gov/TextProposal", nil)
	cdc.RegisterConcrete(&types.DexListProposal{}, "okchain/gov/DexListProposal", nil)
	cdc.RegisterConcrete(&types.ParameterProposal{}, "okchain/gov/ParameterProposal", nil)
	cdc.RegisterConcrete(&types.AppUpgradeProposal{}, "okchain/gov/AppUpgradeProposal", nil)

}



// TODO inefficient BaseResponse unmarshal
func UnmarshalBaseResponse(bz []byte, ptr interface{}) error {
	var br common.BaseResponse
	if err := json.Unmarshal(bz, &br); err != nil {
		return err
	}
	// r.Data is interface{} —— first Marshal then Unmarshal, it's inefficient

	jsonBytes, err := json.Marshal(br.Data)
	if err != nil {
		return err
	}
	// br.Data contains float64 and go-amino doesn't support float
	if err := Cdc.UnmarshalJSON(jsonBytes, ptr); err != nil {
		return err
	}
	return nil
}

// TODO inefficient BaseResponse unmarshal
func UnmarshalJsonBaseResponse(bz []byte, ptr interface{}) error {
	var br common.BaseResponse
	if err := json.Unmarshal(bz, &br); err != nil {
		return err
	}
	// r.Data is interface{} —— first Marshal then Unmarshal, it's inefficient

	jsonBytes, err := json.Marshal(br.Data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonBytes, ptr); err != nil {
		return err
	}
	return nil
}

func UnmarshalListResponse(bz []byte, ptr interface{}) error {
	var lr common.ListResponse
	if err := json.Unmarshal(bz, &lr); err != nil {
		return err
	}

	jsonBytes, err := json.Marshal(lr.Data.Data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, ptr); err != nil {
		return err
	}
	return nil
}
