package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return sdk.MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgSubmitProposal) Route() string                { return "" }
func (MsgSubmitProposal) Type() string                 { return "" }
func (MsgSubmitProposal) ValidateBasic() sdk.Error     { return nil }
func (MsgSubmitProposal) GetSigners() []sdk.AccAddress { return nil }

// MsgDeposit - structure for increasing the deposit on an proposal
type MsgDeposit struct {
	ProposalID uint64         `json:"proposal_id"`
	Depositor  sdk.AccAddress `json:"depositor"`
	Amount     sdk.DecCoins   `json:"amount"`
}

// NewMsgDeposit is a constructor function for MsgDeposit
func NewMsgDeposit(depositor sdk.AccAddress, proposalID uint64, amount sdk.DecCoins) MsgDeposit {
	return MsgDeposit{
		ProposalID: proposalID,
		Depositor:  depositor,
		Amount:     amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDeposit) Route() string                { return "" }
func (MsgDeposit) Type() string                 { return "" }
func (MsgDeposit) ValidateBasic() sdk.Error     { return nil }
func (MsgDeposit) GetSigners() []sdk.AccAddress { return nil }

// MsgVote - structure for voting to an proposal
type MsgVote struct {
	ProposalID uint64         `json:"proposal_id"`
	Voter      sdk.AccAddress `json:"voter"`
	Option     VoteOption     `json:"option"`
}

// NewMsgVote is a constructor function for MsgVote
func NewMsgVote(voter sdk.AccAddress, proposalID uint64, option VoteOption) MsgVote {
	return MsgVote{
		ProposalID: proposalID,
		Voter:      voter,
		Option:     option,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgVote) Route() string                { return "" }
func (MsgVote) Type() string                 { return "" }
func (MsgVote) ValidateBasic() sdk.Error     { return nil }
func (MsgVote) GetSigners() []sdk.AccAddress { return nil }
