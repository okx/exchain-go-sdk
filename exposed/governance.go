package exposed

import (
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
)

// Governance shows the expected behavior for inner governance client
type Governance interface {
	sdk.Module
	GovTx
}

// GovTx shows the expected tx behavior for inner governance client
type GovTx interface {
	SubmitTextProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitParamChangeProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitDelistProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	SubmitCommunityPoolSpendProposal(fromInfo keys.Info, passWd, proposalPath, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Deposit(fromInfo keys.Info, passWd, depositCoinsStr, memo string, proposalID, accNum, seqNum uint64) (sdk.TxResponse, error)
	Vote(fromInfo keys.Info, passWd, voteOption, memo string, proposalID, accNum, seqNum uint64) (sdk.TxResponse, error)
}
