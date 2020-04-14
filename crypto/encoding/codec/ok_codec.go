package codec

// ok codec used specifically for ok client

import (
	"bytes"
	"encoding/json"
	"github.com/okex/okchain-go-sdk/common"
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
	cdc.RegisterInterface((*types.Account)(nil), nil)
	cdc.RegisterConcrete(&types.BaseAccount{}, "cosmos-sdk/Account", nil)
	cdc.RegisterInterface((*types.Proposal)(nil), nil)
	cdc.RegisterConcrete(&types.TextProposal{}, "okchain/gov/TextProposal", nil)
	cdc.RegisterConcrete(&types.DexListProposal{}, "okchain/gov/DexListProposal", nil)
	cdc.RegisterConcrete(&types.ParameterProposal{}, "okchain/gov/ParameterProposal", nil)
	cdc.RegisterConcrete(&types.AppUpgradeProposal{}, "okchain/gov/AppUpgradeProposal", nil)

}

func GetDataFromBaseResponse(bz []byte, ptr interface{}) error {
	dataBytes := getDataFromBaseResponse(bz)
	if err := json.Unmarshal(dataBytes, ptr); err != nil {
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

func getDataFromBaseResponse(bz []byte) []byte {
	preIndex := bytes.Index(bz, []byte("data"))
	//sufIndex := bytes.LastIndex(bz, []byte("detail_msg"))
	return bz[preIndex+6 : len(bz)-1 ]
}
