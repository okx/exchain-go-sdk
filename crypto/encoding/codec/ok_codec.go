package codec

// ok codec used specifically for ok client

import (
	"encoding/json"
	"github.com/ok-chain/ok-gosdk/common"
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

func goSDKRegisterAmino(cdc *amino.Codec) {
	cdc.RegisterInterface((*types.Account)(nil), nil)
	cdc.RegisterConcrete(&types.BaseAccount{}, "auth/Account", nil)

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
	if err := json.Unmarshal(jsonBytes, ptr); err != nil {
		return err
	}
	return nil
}
