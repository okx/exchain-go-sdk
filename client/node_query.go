package client

import (
	"fmt"
	sdktypes "github.com/okex/okchain-go-sdk/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	proposalsInfoPath = "custom/gov/proposals"
	proposalInfoPath  = "custom/gov/proposal"
)

func (cli *OKChainClient) QueryProposals() (sdktypes.Proposals, error) {
	var proposalStatus sdktypes.ProposalStatus
	var voterAddr, depositorAddr sdktypes.AccAddress
	var numLimit uint64
	params := sdktypes.NewQueryProposalsParams(proposalStatus, numLimit, voterAddr, depositorAddr)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryProposalsParams failed in json marshal : %s", err.Error())
	}
	res, err := cli.query(proposalsInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}
	var matchingProposals sdktypes.Proposals
	if err := cli.cdc.UnmarshalJSON(res, &matchingProposals); err != nil {
		return nil, fmt.Errorf("proposals unmarshaled failed : %s", err.Error())
	}

	return matchingProposals, nil
}

func (cli *OKChainClient) QueryProposalByID(proposalID uint64) (sdktypes.Proposal, error) {
	params := sdktypes.NewQueryProposalParams(proposalID)
	jsonBytes, err := cli.cdc.MarshalJSON(params)
	if err != nil {
		return nil, fmt.Errorf("error : QueryProposalParams failed in json marshal : %s", err.Error())
	}
	res, err := cli.query(proposalInfoPath, jsonBytes)
	if err != nil {
		return nil, fmt.Errorf("ok client query error : %s", err.Error())
	}

	var matchingProposal sdktypes.Proposal

	if err := cli.cdc.UnmarshalJSON(res, &matchingProposal); err != nil {
		return nil, fmt.Errorf("proposal unmarshaled failed : %s", err.Error())
	}
	return matchingProposal, nil

}
