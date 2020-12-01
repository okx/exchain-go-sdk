package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)



// MsgEditValidator - structure for editing the info of a validator
type MsgEditValidator struct {
	//Description
	ValidatorAddress sdk.ValAddress `json:"address"`
}

// NewMsgEditValidator creates a msg of edit-validator
func NewMsgEditValidator(valAddr sdk.ValAddress) MsgEditValidator {
	return MsgEditValidator{
		//Description:      description,
		ValidatorAddress: valAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgEditValidator) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgEditValidator) Route() string                { return "" }
func (MsgEditValidator) Type() string                 { return "" }
func (MsgEditValidator) ValidateBasic() sdk.Error     { return nil }
func (MsgEditValidator) GetSigners() []sdk.AccAddress { return nil }

// MsgRegProxy - structure to register delegator as proxy or unregister proxy to delegator
// if Reg == true, action is reg, otherwise action is unreg
type MsgRegProxy struct {
	ProxyAddress sdk.AccAddress `json:"proxy_address"`
	Reg          bool           `json:"reg"`
}

// NewMsgRegProxy creates a msg of registering or unregistering proxy
func NewMsgRegProxy(proxyAddress sdk.AccAddress, reg bool) MsgRegProxy {
	return MsgRegProxy{
		ProxyAddress: proxyAddress,
		Reg:          reg,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgRegProxy) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgRegProxy) Route() string                { return "" }
func (MsgRegProxy) Type() string                 { return "" }
func (MsgRegProxy) ValidateBasic() sdk.Error     { return nil }
func (MsgRegProxy) GetSigners() []sdk.AccAddress { return nil }

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
