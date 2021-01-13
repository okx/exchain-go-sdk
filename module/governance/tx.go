package governance

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/params"
	dexutils "github.com/okex/okexchain/x/dex/client/utils"
	dextypes "github.com/okex/okexchain/x/dex/types"
	distrcli "github.com/okex/okexchain/x/distribution/client/cli"
	distrtypes "github.com/okex/okexchain/x/distribution/types"
	farmutils "github.com/okex/okexchain/x/farm/client/utils"
	farmtypes "github.com/okex/okexchain/x/farm/types"
	govutils "github.com/okex/okexchain/x/gov/client/utils"
	govtypes "github.com/okex/okexchain/x/gov/types"
	paramsutils "github.com/okex/okexchain/x/params/client/utils"
	paramstypes "github.com/okex/okexchain/x/params/types"
)

// SubmitTextProposal submits the text proposal on OKExChain
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

	msg := govtypes.NewMsgSubmitProposal(
		govtypes.ContentFromProposalType(proposal.Title, proposal.Description, proposal.ProposalType),
		deposit,
		fromInfo.GetAddress(),
	)

	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// SubmitParamChangeProposal submits the proposal to change the params on OKExChain
func (gc govClient) SubmitParamsChangeProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := paramsutils.ParseParamChangeProposalJSON(gc.GetCodec(), proposalPath)
	if err != nil {
		return
	}

	msg := govtypes.NewMsgSubmitProposal(
		paramstypes.NewParameterChangeProposal(
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

	proposal, err := dexutils.ParseDelistProposalJSON(gc.GetCodec(), proposalPath)
	if err != nil {
		return
	}

	msg := govtypes.NewMsgSubmitProposal(
		dextypes.NewDelistProposal(
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

// SubmitCommunityPoolSpendProposal submits the proposal to spend the tokens from the community pool on OKExChain
func (gc govClient) SubmitCommunityPoolSpendProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := distrcli.ParseCommunityPoolSpendProposalJSON(gc.GetCodec(), proposalPath)
	if err != nil {
		return
	}

	msg := govtypes.NewMsgSubmitProposal(
		distrtypes.NewCommunityPoolSpendProposal(
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

// SubmitManageWhiteList submits the proposal to manage the white list member of farm module
func (gc govClient) SubmitManageWhiteListProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckKeyParams(fromInfo, passWd); err != nil {
		return
	}

	proposal, err := farmutils.ParseManageWhiteListProposalJSON(gc.GetCodec(), proposalPath)
	if err != nil {
		return
	}

	msg := govtypes.NewMsgSubmitProposal(
		farmtypes.NewManageWhiteListProposal(
			proposal.Title,
			proposal.Description,
			proposal.PoolName,
			proposal.IsAdded,
		),
		proposal.Deposit,
		fromInfo.GetAddress(),
	)

	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Deposit increases the deposit amount on a specific proposal
func (gc govClient) Deposit(fromInfo keys.Info, passWd, depositCoinsStr, memo string, proposalID, accNum,
	seqNum uint64) (resp sdk.TxResponse, err error) {
	if err = params.CheckProposalOperation(fromInfo, passWd, proposalID); err != nil {
		return
	}

	deposit, err := sdk.ParseDecCoins(depositCoinsStr)
	if err != nil {
		return
	}

	msg := govtypes.NewMsgDeposit(fromInfo.GetAddress(), proposalID, deposit)
	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}

// Vote votes for an active proposal
// options: yes/no/no_with_veto/abstain
func (gc govClient) Vote(fromInfo keys.Info, passWd, voteOption, memo string, proposalID, accNum, seqNum uint64) (
	resp sdk.TxResponse, err error) {
	if err = params.CheckProposalOperation(fromInfo, passWd, proposalID); err != nil {
		return
	}

	byteVoteOption, err := govtypes.VoteOptionFromString(govutils.NormalizeVoteOption(voteOption))
	if err != nil {
		return
	}

	msg := govtypes.NewMsgVote(fromInfo.GetAddress(), proposalID, byteVoteOption)
	return gc.BuildAndBroadcast(fromInfo.GetName(), passWd, memo, []sdk.Msg{msg}, accNum, seqNum)
}
