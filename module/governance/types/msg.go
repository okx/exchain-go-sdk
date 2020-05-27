package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// MsgSubmitProposal - structure for submitting the proposal
type MsgSubmitProposal struct {
	Content        Content        `json:"content"`
	InitialDeposit sdk.DecCoins   `json:"initial_deposit"`
	Proposer       sdk.AccAddress `json:"proposer"`
}

// NewMsgSubmitProposal is a constructor function for MsgSubmitProposal
func NewMsgSubmitProposal(content Content, initialDeposit sdk.DecCoins, proposer sdk.AccAddress) MsgSubmitProposal {
	return MsgSubmitProposal{
		Content:        content,
		InitialDeposit: initialDeposit,
		Proposer:       proposer,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgSubmitProposal) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgSubmitProposal) Route() string                { return "" }
func (MsgSubmitProposal) Type() string                 { return "" }
func (MsgSubmitProposal) ValidateBasic() sdk.Error     { return nil }
func (MsgSubmitProposal) GetSigners() []sdk.AccAddress { return nil }
