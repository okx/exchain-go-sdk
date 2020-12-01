package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgBindProxy - structure for binding proxy relationship between voters and voting proxy
type MsgBindProxy struct {
	DelAddr      sdk.AccAddress `json:"delegator_address"`
	ProxyAddress sdk.AccAddress `json:"proxy_address"`
}

// NewMsgBindProxy creates a msg of binding proxy
func NewMsgBindProxy(delAddr, proxyAddr sdk.AccAddress) MsgBindProxy {
	return MsgBindProxy{
		DelAddr:      delAddr,
		ProxyAddress: proxyAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgBindProxy) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgBindProxy) Route() string                { return "" }
func (MsgBindProxy) Type() string                 { return "" }
func (MsgBindProxy) ValidateBasic() sdk.Error     { return nil }
func (MsgBindProxy) GetSigners() []sdk.AccAddress { return nil }

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
