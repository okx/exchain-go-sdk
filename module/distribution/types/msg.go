package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgWithdrawValCommission - structure for validator's reward-withdraw
type MsgWithdrawValCommission struct {
	ValAddr sdk.ValAddress `json:"validator_address"`
}

// NewMsgWithdrawValCommission is a constructor function for MsgWithdrawValCommission
func NewMsgWithdrawValCommission(valAddr sdk.ValAddress) MsgWithdrawValCommission {
	return MsgWithdrawValCommission{
		ValAddr: valAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgWithdrawValCommission) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgWithdrawValCommission) Route() string                { return "" }
func (MsgWithdrawValCommission) Type() string                 { return "" }
func (MsgWithdrawValCommission) ValidateBasic() sdk.Error     { return nil }
func (MsgWithdrawValCommission) GetSigners() []sdk.AccAddress { return nil }
