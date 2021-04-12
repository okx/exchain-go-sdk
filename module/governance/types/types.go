package types

import (
	govtypes "github.com/okex/exchain/x/gov/types"
)

// const
const (
	ModuleName = govtypes.ModuleName
)

type (
	Proposal = govtypes.Proposal
)

// ProposalJSON - structure for a standard proposal from the JSON file
type ProposalJSON struct {
	Title        string
	Description  string
	ProposalType string
	Deposit      string
}
