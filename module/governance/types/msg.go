package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
