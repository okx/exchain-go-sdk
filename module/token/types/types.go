package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// const
const (
	ModuleName = "token"

	AccountTokensInfoPath = "custom/token/accounts"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for token module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgSend{}, "okexchain/token/MsgTransfer")
	cdc.RegisterConcrete(MsgMultiSend{}, "okexchain/token/MsgMultiTransfer")
	cdc.RegisterConcrete(MsgTokenIssue{}, "okexchain/token/MsgIssue")
	cdc.RegisterConcrete(MsgTokenMint{}, "okexchain/token/MsgMint")
	cdc.RegisterConcrete(MsgTokenBurn{}, "okexchain/token/MsgBurn")
	cdc.RegisterConcrete(MsgTokenModify{}, "okexchain/token/MsgModify")
}

// TransferUnit - amount part for multi-send
type TransferUnit struct {
	To    sdk.AccAddress `json:"to"`
	Coins sdk.DecCoins   `json:"coins"`
}

// NewTransferUnit creates a new instance of TransferUnit
func NewTransferUnit(addr sdk.AccAddress, coins sdk.DecCoins) TransferUnit {
	return TransferUnit{
		To:    addr,
		Coins: coins,
	}
}

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
