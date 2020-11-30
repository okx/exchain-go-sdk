package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
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
)

var (
	msgCdc = codec.New()
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

// AccountTokensInfo - structure for available tokens info of an account
type AccountTokensInfo struct {
	Address    string     `json:"address"`
	Currencies []CoinInfo `json:"currencies"`
}

// CoinInfo - structure for a kind of currencies in AccountTokensInfo
type CoinInfo struct {
	Symbol    string `json:"symbol"`
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
	Locked    string `json:"locked"`
}
