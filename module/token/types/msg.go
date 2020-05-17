package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// MsgSend - structure to transfer
type MsgSend struct {
	FromAddress sdk.AccAddress `json:"from_address"`
	ToAddress   sdk.AccAddress `json:"to_address"`
	Amount      sdk.DecCoins   `json:"amount"`
}

// NewMsgTokenSend is a constructor function for MsgSend
func NewMsgTokenSend(fromAddr, toAddr sdk.AccAddress, coins sdk.DecCoins) MsgSend {
	return MsgSend{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Amount:      coins,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgSend) Route() string                { return "" }
func (MsgSend) Type() string                 { return "" }
func (MsgSend) ValidateBasic() sdk.Error     { return nil }
func (MsgSend) GetSigners() []sdk.AccAddress { return nil }

// MsgMultiSend - structure to transfer to multi receivers
type MsgMultiSend struct {
	From      sdk.AccAddress `json:"from"`
	Transfers []TransferUnit `json:"transfers"`
}

// NewMsgMultiSend is a constructor function for MsgMultiSend
func NewMsgMultiSend(fromAddr sdk.AccAddress, transfers []TransferUnit) MsgMultiSend {
	return MsgMultiSend{
		From:      fromAddr,
		Transfers: transfers,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgMultiSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgMultiSend) Route() string                { return "" }
func (MsgMultiSend) Type() string                 { return "" }
func (MsgMultiSend) ValidateBasic() sdk.Error     { return nil }
func (MsgMultiSend) GetSigners() []sdk.AccAddress { return nil }

type MsgTokenBurn struct {
	Amount sdk.DecCoin    `json:"amount"`
	Owner  sdk.AccAddress `json:"owner"`
}

func NewMsgTokenBurn(amount sdk.DecCoin, owner sdk.AccAddress) MsgTokenBurn {
	return MsgTokenBurn{
		Amount: amount,
		Owner:  owner,
	}
}

func (msg MsgTokenBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

func (msg MsgTokenBurn) Route() string { return "" }
func (msg MsgTokenBurn) Type() string { return "" }
func (msg MsgTokenBurn) ValidateBasic() sdk.Error {return nil}
func (msg MsgTokenBurn) GetSigners() []sdk.AccAddress {return nil}

type MsgTokenMint struct {
	Amount sdk.DecCoin    `json:"amount"`
	Owner  sdk.AccAddress `json:"owner"`
}

func NewMsgTokenMint(amount sdk.DecCoin, owner sdk.AccAddress) MsgTokenMint {
	return MsgTokenMint{
		Amount: amount,
		Owner:  owner,
	}
}

func (msg MsgTokenMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

func (msg MsgTokenMint) Route() string { return "" }
func (msg MsgTokenMint) Type() string { return "" }
func (msg MsgTokenMint) ValidateBasic() sdk.Error { return nil}
func (msg MsgTokenMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgTokenIssue - structure to issue a kind of token
type MsgTokenIssue struct {
	Description    string         `json:"description"`
	Symbol         string         `json:"symbol"`
	OriginalSymbol string         `json:"original_symbol"`
	WholeName      string         `json:"whole_name"`
	TotalSupply    string         `json:"total_supply"`
	Owner          sdk.AccAddress `json:"owner"`
	Mintable       bool           `json:"mintable"`
}

// NewMsgTokenIssue creates a new instance of MsgTokenIssue
func NewMsgTokenIssue(owner sdk.AccAddress, tokenDesc, symbol, originalSymbol, wholeName, totalSupply string,
	mintable bool) MsgTokenIssue {
	return MsgTokenIssue{
		Description:    tokenDesc,
		Symbol:         symbol,
		OriginalSymbol: originalSymbol,
		WholeName:      wholeName,
		TotalSupply:    totalSupply,
		Owner:          owner,
		Mintable:       mintable,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgTokenIssue) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgTokenIssue) Route() string                { return "" }
func (MsgTokenIssue) Type() string                 { return "" }
func (MsgTokenIssue) ValidateBasic() sdk.Error     { return nil }
func (MsgTokenIssue) GetSigners() []sdk.AccAddress { return nil }
