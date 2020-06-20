package types

import (
	"encoding/json"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

// MsgDeposit - structure for depositing to the delegator account
type MsgDeposit struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address"`
	Amount           sdk.DecCoin    `json:"quantity"`
}

// NewMsgDeposit creates a msg of delegating
func NewMsgDeposit(delAddr sdk.AccAddress, amount sdk.DecCoin) MsgDeposit {
	return MsgDeposit{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDeposit) Route() string                { return "" }
func (MsgDeposit) Type() string                 { return "" }
func (MsgDeposit) ValidateBasic() sdk.Error     { return nil }
func (MsgDeposit) GetSigners() []sdk.AccAddress { return nil }

// MsgUndelegate - structure for delegating to exchange the votes
type MsgUndelegate struct {
	DelegatorAddress sdk.AccAddress `json:"delegator_address" `
	Amount           sdk.DecCoin    `json:"quantity"`
}

// NewMsgUndelegate creates a msg of undelegating
func NewMsgUndelegate(delAddr sdk.AccAddress, amount sdk.DecCoin) MsgUndelegate {
	return MsgUndelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUndelegate) Route() string                { return "" }
func (MsgUndelegate) Type() string                 { return "" }
func (MsgUndelegate) ValidateBasic() sdk.Error     { return nil }
func (MsgUndelegate) GetSigners() []sdk.AccAddress { return nil }

// MsgVote - structure for voting transactions
type MsgVote struct {
	DelAddr  sdk.AccAddress   `json:"delegator_address"`
	ValAddrs []sdk.ValAddress `json:"validator_addresses"`
}

// NewMsgVote creates a msg of multi voting
func NewMsgVote(delAddr sdk.AccAddress, valAddrs []sdk.ValAddress) MsgVote {
	return MsgVote{
		DelAddr:  delAddr,
		ValAddrs: valAddrs,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgVote) Route() string                { return "" }
func (MsgVote) Type() string                 { return "" }
func (MsgVote) ValidateBasic() sdk.Error     { return nil }
func (MsgVote) GetSigners() []sdk.AccAddress { return nil }

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
		Description:       msg.Description,
		DelegatorAddress:  msg.DelegatorAddress,
		ValidatorAddress:  msg.ValidatorAddress,
		PubKey:            sdk.MustBech32ifyConsPub(msg.PubKey),
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
