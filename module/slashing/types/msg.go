package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// MsgUnjail - structure to recover a jailed validator
type MsgUnjail struct {
	ValidatorAddr sdk.ValAddress `json:"address"`
}

// NewMsgUnjail creates a msg of unjailing
func NewMsgUnjail(validatorAddr sdk.ValAddress) MsgUnjail {
	return MsgUnjail{
		ValidatorAddr: validatorAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnjail) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnjail) Route() string                { return "" }
func (MsgUnjail) Type() string                 { return "" }
func (MsgUnjail) ValidateBasic() sdk.Error     { return nil }
func (MsgUnjail) GetSigners() []sdk.AccAddress { return nil }
