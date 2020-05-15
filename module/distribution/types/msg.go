package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// MsgSetWithdrawAddr - structure for changing the reward-withdraw address
type MsgSetWithdrawAddr struct {
	DelAddr      sdk.AccAddress `json:"delegator_address"`
	WithdrawAddr sdk.AccAddress `json:"withdraw_address"`
}

// NewMsgSetWithdrawAddr is a constructor function for MsgSetWithdrawAddr
func NewMsgSetWithdrawAddr(delAddr, withdrawAddr sdk.AccAddress) MsgSetWithdrawAddr {
	return MsgSetWithdrawAddr{
		DelAddr:      delAddr,
		WithdrawAddr: withdrawAddr,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgSetWithdrawAddr) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgSetWithdrawAddr) Route() string                { return "" }
func (MsgSetWithdrawAddr) Type() string                 { return "" }
func (MsgSetWithdrawAddr) ValidateBasic() sdk.Error     { return nil }
func (MsgSetWithdrawAddr) GetSigners() []sdk.AccAddress { return nil }
