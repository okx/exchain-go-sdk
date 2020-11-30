package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain/x/token"
	tokentypes "github.com/okex/okexchain/x/token/types"
)

// const
const (
	ModuleName = token.ModuleName

	AccountTokensInfoPath = "custom/token/accounts"
)

type (
	TransferUnit = tokentypes.TransferUnit
	TokenResp    = tokentypes.TokenResp
)

// Token - structure for detail info of a kind of token
type Token struct {
	Description         string         `json:"description"`
	Symbol              string         `json:"symbol"`
	OriginalSymbol      string         `json:"original_symbol"`
	WholeName           string         `json:"whole_name"`
	OriginalTotalSupply sdk.Dec        `json:"original_total_supply"`
	TotalSupply         sdk.Dec        `json:"total_supply"`
	Owner               sdk.AccAddress `json:"owner"`
	Mintable            bool           `json:"mintable"`
}
