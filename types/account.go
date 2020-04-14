package types

type AccountTokensInfo struct {
	Address    string    `json:"address"`
	Currencies CoinsInfo `json:"currencies"`
}
