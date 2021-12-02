package exposed

import (
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain-go-sdk/module/governance/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
)

// Governance shows the expected behavior for inner governance client
type Governance interface {
	gosdktypes.Module
	GovTx
	GovQuery
}

// GovTx shows the expected tx behavior for inner governance client
type GovTx interface {
	SubmitTextProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitParamsChangeProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitDelistProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitCommunityPoolSpendProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitManageWhiteListProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Deposit(fromInfo keys.Info, passWd, depositCoinsStr, memo string, proposalID, accNum, seqNum uint64) (sdk.TxResponse, error)
	Vote(fromInfo keys.Info, passWd, voteOption, memo string, proposalID, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// GovQuery shows the expected query behavior for inner governance client
type GovQuery interface {
	QueryProposals(depositorAddrStr, voterAddrStr, status string, numLimit uint64) ([]types.Proposal, error)
}
