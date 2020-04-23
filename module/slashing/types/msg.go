package types

import "github.com/okex/okchain-go-sdk/types"

// MsgUnjail - structure to recover a jailed validator
type MsgUnjail struct {
	ValidatorAddr types.ValAddress `json:"address"`
}

// NewMsgUnjail creates a msg of unjailing
func NewMsgUnjail(validatorAddr types.ValAddress) MsgUnjail {
	return MsgUnjail{
		ValidatorAddr: validatorAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnjail) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnjail) Route() string                  { return "" }
func (MsgUnjail) Type() string                   { return "" }
func (MsgUnjail) ValidateBasic() types.Error     { return nil }
func (MsgUnjail) GetSigners() []types.AccAddress { return nil }
