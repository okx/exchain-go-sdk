package types

type AccountTokensInfo struct {
	Address    string    `json:"address"`
	Currencies CoinsInfo `json:"currencies"`
}

type CoinsInfo []CoinInfo

type CoinInfo struct {
	Symbol    string `json:"symbol"`
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
	Locked    string `json:"locked"`
}