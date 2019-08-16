package types

type CoinsInfo []CoinInfo

type CoinInfo struct {
	Symbol    string `json:"symbol"`
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
	Locked    string `json:"locked"`
}

type Token struct {
	Desc           string     `json:"desc"`
	Symbol         string     `json:"symbol"`
	OriginalSymbol string     `json:"originalSymbol"`
	WholeName      string     `json:"wholeName"`
	TotalSupply    int64      `json:"totalSupply"`
	Owner          AccAddress `json:"owner"`
	Mintable       bool       `json:"mintable"`
}

type TokenPair struct {
	BaseAssetSymbol  string `json:"baseAssetSymbol"`
	QuoteAssetSymbol string `json:"quoteAssetSymbol"`
	InitPrice        Dec    `json:"price"`
	MaxPriceDigit    int64  `json:"maxPriceDigit"`
	MaxQuantityDigit int64  `json:"maxSizeDigit"`
	MinQuantity      Dec    `json:"minTradeSize"`
}
