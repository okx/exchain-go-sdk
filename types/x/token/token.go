package token

type CoinInfo struct {
	Symbol    string `json:"symbol"`
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
	Locked    string `json:"locked"`
}

type CoinsInfo []CoinInfo

type AccountResponse struct {
	Address    string    `json:"address"`
	Currencies CoinsInfo `json:"currencies"`
}