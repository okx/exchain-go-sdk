package governance

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/governance/types"
	"github.com/okex/okexchain-go-sdk/utils"
	govutils "github.com/okex/okexchain/x/gov/client/utils"
	govtypes "github.com/okex/okexchain/x/gov/types"
)

// QueryProposals gets all proposals
// Note:
//	optional:
//		status option - DepositPeriod|VotingPeriod|Passed|Rejected. Defaults to all proposals by ""
//		depositorAddrStr - filter by proposals deposited on by depositor. Defaults to all proposals by ""
//		voterAddrStr - filter by proposals voted on by voted. Defaults to all proposals by ""
// 		numLimit - limit to latest [number] proposals. Defaults to all proposals by 0
func (gc govClient) QueryProposals(depositorAddrStr, voterAddrStr, status string, numLimit uint64) (
	proposals []types.Proposal, err error) {
	var depositorAddr, voterAddr sdk.AccAddress
	var proposalStatus govtypes.ProposalStatus

	proposalParams := govtypes.NewQueryProposalsParams(proposalStatus, numLimit, depositorAddr, voterAddr)
	if len(depositorAddrStr) != 0 {
		depositorAddr, err = sdk.AccAddressFromBech32(depositorAddrStr)
		if err != nil {
			return
		}
		proposalParams.Depositor = depositorAddr
	}

	if len(voterAddrStr) != 0 {
		voterAddr, err = sdk.AccAddressFromBech32(voterAddrStr)
		if err != nil {
			return
		}
		proposalParams.Voter = voterAddr
	}

	if len(status) != 0 {
		proposalStatus, err = govtypes.ProposalStatusFromString(govutils.NormalizeProposalStatus(status))
		if err != nil {
			return
		}
		proposalParams.ProposalStatus = proposalStatus
	}

	jsonBytes, err := gc.GetCodec().MarshalJSON(proposalParams)
	if err != nil {
		return proposals, utils.ErrMarshalJSON(err.Error())
	}

	path := fmt.Sprintf("custom/%s/proposals", govtypes.QuerierRoute)
	res, _, err := gc.Query(path, jsonBytes)
	if err != nil {
		return proposals, utils.ErrClientQuery(err.Error())
	}

	if err = gc.GetCodec().UnmarshalJSON(res, &proposals); err != nil {
		return proposals, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
