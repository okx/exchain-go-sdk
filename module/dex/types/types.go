package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain/x/common"
	"github.com/okex/okexchain/x/dex"
	dextypes "github.com/okex/okexchain/x/dex/types"
)

// const
const (
	ModuleName = dextypes.ModuleName
)

var (
	msgCdc = gosdktypes.NewCodec()
)

func init() {
	gosdktypes.RegisterBasicCodec(msgCdc)
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for dex module
func RegisterCodec(cdc *codec.Codec) {
	dex.RegisterCodec(cdc)
}

type (
	TokenPair = dextypes.TokenPair
)

// ListResponse - used for decoding
type ListResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detail_msg"`
	Data      ListDataRes `json:"data"`
}

// ListDataRes - used for decoding
type ListDataRes struct {
	Data      []TokenPair      `json:"data"`
	ParamPage common.ParamPage `json:"param_page"`
}
