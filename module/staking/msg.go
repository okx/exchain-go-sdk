package staking

import (
	"encoding/json"
	"github.com/okex/okchain-go-sdk/exposed"
	"github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/tendermint/tendermint/crypto"
)

type MsgDelegate struct {
	DelegatorAddress types.AccAddress `json:"delegator_address"`
	Amount           types.DecCoin    `json:"quantity"`
}

// NewMsgDelegate creates a msg of delegating
func NewMsgDelegate(delAddr types.AccAddress, amount types.DecCoin) MsgDelegate {
	return MsgDelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDelegate) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDelegate) Route() string                  { return "" }
func (MsgDelegate) Type() string                   { return "" }
func (MsgDelegate) ValidateBasic() types.Error     { return nil }
func (MsgDelegate) GetSigners() []types.AccAddress { return nil }

type MsgUndelegate struct {
	DelegatorAddress types.AccAddress `json:"delegator_address" `
	Amount           types.DecCoin    `json:"quantity"`
}

// NewMsgUndelegate creates a msg of undelegating
func NewMsgUndelegate(delAddr types.AccAddress, amount types.DecCoin) MsgUndelegate {
	return MsgUndelegate{
		DelegatorAddress: delAddr,
		Amount:           amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUndelegate) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUndelegate) Route() string                  { return "" }
func (MsgUndelegate) Type() string                   { return "" }
func (MsgUndelegate) ValidateBasic() types.Error     { return nil }
func (MsgUndelegate) GetSigners() []types.AccAddress { return nil }

type MsgVote struct {
	DelAddr  types.AccAddress   `json:"delegator_address"`
	ValAddrs []types.ValAddress `json:"validator_addresses"`
}

// NewMsgVote creates a msg of multi voting
func NewMsgVote(delAddr types.AccAddress, valAddrs []types.ValAddress) MsgVote {
	return MsgVote{
		DelAddr:  delAddr,
		ValAddrs: valAddrs,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgVote) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgVote) Route() string                  { return "" }
func (MsgVote) Type() string                   { return "" }
func (MsgVote) ValidateBasic() types.Error     { return nil }
func (MsgVote) GetSigners() []types.AccAddress { return nil }

type MsgDestroyValidator struct {
	DelAddr types.AccAddress `json:"delegator_address"`
}

// NewMsgDestroyValidator creates a msg of destroy-validator
func NewMsgDestroyValidator(delAddr types.AccAddress) MsgDestroyValidator {
	return MsgDestroyValidator{
		DelAddr: delAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDestroyValidator) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDestroyValidator) Route() string                  { return "" }
func (MsgDestroyValidator) Type() string                   { return "" }
func (MsgDestroyValidator) ValidateBasic() types.Error     { return nil }
func (MsgDestroyValidator) GetSigners() []types.AccAddress { return nil }

type MsgCreateValidator struct {
	Description       exposed.Description `json:"description"`
	MinSelfDelegation types.DecCoin       `json:"min_self_delegation"`
	DelegatorAddress  types.AccAddress    `json:"delegator_address"`
	ValidatorAddress  types.ValAddress    `json:"validator_address"`
	PubKey            crypto.PubKey       `json:"pubkey"`
}

type msgCreateValidatorJSON struct {
	Description       exposed.Description `json:"description"`
	MinSelfDelegation types.DecCoin       `json:"min_self_delegation"`
	DelegatorAddress  types.AccAddress    `json:"delegator_address"`
	ValidatorAddress  types.ValAddress    `json:"validator_address"`
	PubKey            string              `json:"pubkey"`
}

// NewMsgCreateValidator creates a msg of create-validator
// Delegator address and validator address are the same
func NewMsgCreateValidator(valAddr types.ValAddress, pubKey crypto.PubKey, description exposed.Description,
) MsgCreateValidator {
	minSelfDelegationCoin, err := utils.ParseDecCoin(defaultMinSelfDelegation)
	if err != nil {
		panic(err)
	}
	return MsgCreateValidator{
		Description:       description,
		DelegatorAddress:  types.AccAddress(valAddr),
		ValidatorAddress:  valAddr,
		PubKey:            pubKey,
		MinSelfDelegation: minSelfDelegationCoin,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateValidator) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// MarshalJSON is useful for the signing of msg MsgCreateValidator
func (msg MsgCreateValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(msgCreateValidatorJSON{
		Description:       msg.Description,
		DelegatorAddress:  msg.DelegatorAddress,
		ValidatorAddress:  msg.ValidatorAddress,
		PubKey:            types.MustBech32ifyConsPub(msg.PubKey),
		MinSelfDelegation: msg.MinSelfDelegation,
	})
}

// nolint
func (MsgCreateValidator) Route() string                  { return "" }
func (MsgCreateValidator) Type() string                   { return "" }
func (MsgCreateValidator) ValidateBasic() types.Error     { return nil }
func (MsgCreateValidator) GetSigners() []types.AccAddress { return nil }

type MsgEditValidator struct {
	exposed.Description
	ValidatorAddress types.ValAddress `json:"address"`
}

// NewMsgEditValidator creates a msg of edit-validator
func NewMsgEditValidator(valAddr types.ValAddress, description exposed.Description) MsgEditValidator {
	return MsgEditValidator{
		Description:      description,
		ValidatorAddress: valAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgEditValidator) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgEditValidator) Route() string                  { return "" }
func (MsgEditValidator) Type() string                   { return "" }
func (MsgEditValidator) ValidateBasic() types.Error     { return nil }
func (MsgEditValidator) GetSigners() []types.AccAddress { return nil }

type MsgRegProxy struct {
	ProxyAddress types.AccAddress `json:"proxy_address"`
	Reg          bool             `json:"reg"`
}

// NewMsgRegProxy creates a msg of registering or unregistering proxy
func NewMsgRegProxy(proxyAddress types.AccAddress, reg bool) MsgRegProxy {
	return MsgRegProxy{
		ProxyAddress: proxyAddress,
		Reg:          reg,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgRegProxy) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgRegProxy) Route() string                  { return "" }
func (MsgRegProxy) Type() string                   { return "" }
func (MsgRegProxy) ValidateBasic() types.Error     { return nil }
func (MsgRegProxy) GetSigners() []types.AccAddress { return nil }

type MsgBindProxy struct {
	DelAddr      types.AccAddress `json:"delegator_address"`
	ProxyAddress types.AccAddress `json:"proxy_address"`
}

// NewMsgBindProxy creates a msg of binding proxy
func NewMsgBindProxy(delAddr, proxyAddr types.AccAddress) MsgBindProxy {
	return MsgBindProxy{
		DelAddr:      delAddr,
		ProxyAddress: proxyAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgBindProxy) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgBindProxy) Route() string                  { return "" }
func (MsgBindProxy) Type() string                   { return "" }
func (MsgBindProxy) ValidateBasic() types.Error     { return nil }
func (MsgBindProxy) GetSigners() []types.AccAddress { return nil }

type MsgUnbindProxy struct {
	DelAddr types.AccAddress `json:"delegator_address"`
}

// NewMsgUnbindProxy creates a msg of unbinding proxy
func NewMsgUnbindProxy(delAddr types.AccAddress) MsgUnbindProxy {
	return MsgUnbindProxy{
		DelAddr: delAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnbindProxy) GetSignBytes() []byte {
	return types.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnbindProxy) Route() string                  { return "" }
func (MsgUnbindProxy) Type() string                   { return "" }
func (MsgUnbindProxy) ValidateBasic() types.Error     { return nil }
func (MsgUnbindProxy) GetSigners() []types.AccAddress { return nil }
