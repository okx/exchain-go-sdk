package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// MsgDestroyValidator - structure to deregister a validator
type MsgDestroyValidator struct {
	DelAddr sdk.AccAddress `json:"delegator_address"`
}

// NewMsgDestroyValidator creates a msg of destroy-validator
func NewMsgDestroyValidator(delAddr sdk.AccAddress) MsgDestroyValidator {
	return MsgDestroyValidator{
		DelAddr: delAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDestroyValidator) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDestroyValidator) Route() string                { return "" }
func (MsgDestroyValidator) Type() string                 { return "" }
func (MsgDestroyValidator) ValidateBasic() sdk.Error     { return nil }
func (MsgDestroyValidator) GetSigners() []sdk.AccAddress { return nil }

// MsgCreateValidator - structure for creating a validator
type MsgCreateValidator struct {
	Description       Description    `json:"description"`
	MinSelfDelegation sdk.DecCoin    `json:"min_self_delegation"`
	DelegatorAddress  sdk.AccAddress `json:"delegator_address"`
	ValidatorAddress  sdk.ValAddress `json:"validator_address"`
	PubKey            crypto.PubKey  `json:"pubkey"`
}

type msgCreateValidatorJSON struct {
	Description       Description    `json:"description"`
	MinSelfDelegation sdk.DecCoin    `json:"min_self_delegation"`
	DelegatorAddress  sdk.AccAddress `json:"delegator_address"`
	ValidatorAddress  sdk.ValAddress `json:"validator_address"`
	PubKey            string         `json:"pubkey"`
}

// NewMsgCreateValidator creates a msg of create-validator
// Delegator address and validator address are the same
func NewMsgCreateValidator(valAddr sdk.ValAddress, pubKey crypto.PubKey, description Description,
) MsgCreateValidator {
	minSelfDelegationCoin, err := sdk.ParseDecCoin(defaultMinSelfDelegation)
	if err != nil {
		panic(err)
	}
	return MsgCreateValidator{
		Description:       description,
		DelegatorAddress:  sdk.AccAddress(valAddr),
		ValidatorAddress:  valAddr,
		PubKey:            pubKey,
		MinSelfDelegation: minSelfDelegationCoin,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateValidator) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// MarshalJSON is useful for the signing of msg MsgCreateValidator
func (msg MsgCreateValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Description:      msg.Description,
		DelegatorAddress: msg.DelegatorAddress,
		ValidatorAddress: msg.ValidatorAddress,
		// TODO
		//PubKey:            sdk.MustBech32ifyConsPub(msg.PubKey),
		PubKey:            sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeConsPub, msg.PubKey),
		MinSelfDelegation: msg.MinSelfDelegation,
	})
}

// nolint
func (MsgCreateValidator) Route() string                { return "" }
func (MsgCreateValidator) Type() string                 { return "" }
func (MsgCreateValidator) ValidateBasic() sdk.Error     { return nil }
func (MsgCreateValidator) GetSigners() []sdk.AccAddress { return nil }

// MsgEditValidator - structure for editing the info of a validator
type MsgEditValidator struct {
	Description
	ValidatorAddress sdk.ValAddress `json:"address"`
}

// NewMsgEditValidator creates a msg of edit-validator
func NewMsgEditValidator(valAddr sdk.ValAddress, description Description) MsgEditValidator {
	return MsgEditValidator{
		Description:      description,
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
