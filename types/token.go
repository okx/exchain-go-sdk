package types

import "encoding/json"

type CoinsInfo []CoinInfo

type CoinInfo struct {
	Symbol    string `json:"symbol"`
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
	Locked    string `json:"locked"`
}

type Token struct {
	Desc           string     `json:"description"`
	Symbol         string     `json:"symbol"`
	OriginalSymbol string     `json:"original_symbol"`
	WholeName      string     `json:"whole_name"`
	TotalSupply    Dec        `json:"total_supply"`
	Owner          AccAddress `json:"owner"`
	Mintable       bool       `json:"mintable"`
}

func (token Token) String() string {
	b, _ := json.Marshal(token)
	return string(b)
}

type TokenPair struct {
	BaseAssetSymbol  string `json:"base_asset_symbol"`
	QuoteAssetSymbol string `json:"quote_asset_symbol"`
	InitPrice        Dec    `json:"price"`
	MaxPriceDigit    int64  `json:"max_price_digit"`
	MaxQuantityDigit int64  `json:"max_size_digit"`
	MinQuantity      Dec    `json:"min_trade_size"`
	TokenPairId      string `json:"token_pair_id"`
}

type TransferUnit struct {
	To    AccAddress `json:"to"`
	Coins DecCoins   `json:"coins"`
}

func NewTransferUnit(addr AccAddress, coins DecCoins) TransferUnit {
	return TransferUnit{
		addr,
		coins,
	}
}
