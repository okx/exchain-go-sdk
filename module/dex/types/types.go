package types

import (
	"github.com/okex/exchain/x/common"
	dextypes "github.com/okex/exchain/x/dex/types"
)

// const
const (
	ModuleName = dextypes.ModuleName
)

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
