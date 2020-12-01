package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgUnbindProxy - structure for unbinding proxy relationship between voters and proxy
type MsgUnbindProxy struct {
	DelAddr sdk.AccAddress `json:"delegator_address"`
}

// NewMsgUnbindProxy creates a msg of unbinding proxy
func NewMsgUnbindProxy(delAddr sdk.AccAddress) MsgUnbindProxy {
	return MsgUnbindProxy{
		DelAddr: delAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnbindProxy) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnbindProxy) Route() string                { return "" }
func (MsgUnbindProxy) Type() string                 { return "" }
func (MsgUnbindProxy) ValidateBasic() sdk.Error     { return nil }
func (MsgUnbindProxy) GetSigners() []sdk.AccAddress { return nil }
