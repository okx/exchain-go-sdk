package governance

import (
	"github.com/okex/okchain-go-sdk/module/governance/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
	"github.com/okex/okchain-go-sdk/types/params"
)

// SubmitTextProposal submits the text proposal on OKChain
func (gc govClient) SubmitTextProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := parseProposalFromFile(proposalPath)
	if err != nil {
		return
	}

	deposit, err := sdk.ParseDecCoins(proposal.Deposit)
	if err != nil {
		return
	}

	msg := types.NewMsgSubmitProposal(
		types.NewTextProposal(proposal.Title, proposal.Description),
		deposit,
		fromInfo.GetAddress(),
	)

	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// SubmitParamChangeProposal submits the proposal to change the params on OKChain
func (gc govClient) SubmitParamChangeProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := parseParamChangeProposalFromFile(proposalPath)
	if err != nil {
		return
	}

	msg := types.NewMsgSubmitProposal(
		types.NewParameterChangeProposal(
			proposal.Title,
			proposal.Description,
			proposal.Changes.ToParamChanges(),
			proposal.Height,
		),
		proposal.Deposit,
		fromInfo.GetAddress(),
	)

	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// SubmitDelistProposal submits the proposal to delist a token pair from dex
func (gc govClient) SubmitDelistProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := parseDelistProposalFromFile(proposalPath)
	if err != nil {
		return
	}

	msg := types.NewMsgSubmitProposal(
		types.NewDelistProposal(
			proposal.Title,
			proposal.Description,
			fromInfo.GetAddress(),
			proposal.BaseAsset,
			proposal.QuoteAsset,
		),
		proposal.Deposit,
		fromInfo.GetAddress(),
	)

	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}

// SubmitCommunityPoolSpendProposal submits the proposal to spend the tokens from the community pool on OKChain
func (gc govClient) SubmitCommunityPoolSpendProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := parseCommunityPoolSpendProposalFromFile(proposalPath)
	if err != nil {
		return
	}

	msg := types.NewMsgSubmitProposal(
		types.NewCommunityPoolSpendProposal(
			proposal.Title,
			proposal.Description,
			proposal.Recipient,
			proposal.Amount,
		),
		proposal.Deposit,
		fromInfo.GetAddress(),
	)

	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)

}
