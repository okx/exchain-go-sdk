package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// const
const (
	ModuleName = "token"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for token module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgSend{}, "okchain/token/MsgTransfer")
	cdc.RegisterConcrete(MsgMultiSend{}, "okchain/token/MsgMultiTransfer")
	cdc.RegisterConcrete(MsgTokenIssue{}, "okchain/token/MsgIssue")
	cdc.RegisterConcrete(MsgMint{}, "okchain/token/MsgMint")
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
