package types

type DecCoins []DecCoin

type DecCoin struct {
	Denom  string `json:"denom"`
	Amount Dec    `json:"amount"`
}