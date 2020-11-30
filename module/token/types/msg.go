package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgTokenBurn - structure to burn token
type MsgTokenBurn struct {
	Amount sdk.DecCoin    `json:"amount"`
	Owner  sdk.AccAddress `json:"owner"`
}

// NewMsgTokenBurn is a constructor function for MsgTokenBurn
func NewMsgTokenBurn(amount sdk.DecCoin, owner sdk.AccAddress) MsgTokenBurn {
	return MsgTokenBurn{
		Amount: amount,
		Owner:  owner,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgTokenBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (msg MsgTokenBurn) Route() string                { return "" }
func (msg MsgTokenBurn) Type() string                 { return "" }
func (msg MsgTokenBurn) ValidateBasic() sdk.Error     { return nil }
func (msg MsgTokenBurn) GetSigners() []sdk.AccAddress { return nil }

// MsgTokenModify - structure to edit the info of a specific token
type MsgTokenModify struct {
	Owner                 sdk.AccAddress `json:"owner"`
	Symbol                string         `json:"symbol"`
	Description           string         `json:"description"`
	WholeName             string         `json:"whole_name"`
	IsDescriptionModified bool           `json:"description_modified"`
	IsWholeNameModified   bool           `json:"whole_name_modified"`
}

// NewMsgTokenModify creates a new instance of MsgTokenModify
func NewMsgTokenModify(symbol, desc, wholeName string, isDescEdit, isWholeNameEdit bool, owner sdk.AccAddress) MsgTokenModify {
	return MsgTokenModify{
		Symbol:                symbol,
		IsDescriptionModified: isDescEdit,
		Description:           desc,
		IsWholeNameModified:   isWholeNameEdit,
		WholeName:             wholeName,
		Owner:                 owner,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgTokenModify) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgTokenModify) Route() string                { return "" }
func (MsgTokenModify) Type() string                 { return "" }
func (MsgTokenModify) ValidateBasic() sdk.Error     { return nil }
func (MsgTokenModify) GetSigners() []sdk.AccAddress { return nil }
