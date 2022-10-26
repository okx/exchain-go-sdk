package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"github.com/okex/exchain/x/feesplit/types"
)

// Feesplit shows the expected behavior for inner Feesplit client
type Feesplit interface {
	gosdktypes.Module
	FeesplitTx
	FeesplitQuery
}

// FeesplitTx shows the expected tx behavior for inner Feesplit client
type FeesplitTx interface {
	// RegisterFeeSplit register a contract for fee distribution.
	// **NOTE** Please ensure, that the deployer of the contract (or the factory that deployes the contract)
	// is an account that is owned by your project
	// to avoid that an individual deployer who leaves your project becomes malicious.
	RegisterFeeSplit(fromInfo keys.Info, passWd string, accNum, seqNum uint64, memo string, contractAddress string, nonces []uint64, withdrawAddress string) (*sdk.TxResponse, error)

	// CancelFeeSplit cancel a contract from fee distribution.
	// The deployer will no longer receive fees from users interacting with the contract.
	// Only the contract deployer can cancel a contract.
	CancelFeeSplit(fromInfo keys.Info, passWd string, accNum, seqNum uint64, memo string, contractAddress string) (*sdk.TxResponse, error)

	// UpdateFeeSplit update withdraw address for a contract registered for fee distribution.
	// Only the contract deployer can update the withdraw address.
	UpdateFeeSplit(fromInfo keys.Info, passWd string, accNum, seqNum uint64, memo string, contractAddress string, withdrawAddress string) (*sdk.TxResponse, error)

	FeeSplitSharesProposal(fromInfo keys.Info, passWd string, accNum, seqNum uint64, memo string)
}

// FeesplitQuery shows the expected query behavior for inner Feesplit client
type FeesplitQuery interface {

	// QueryFeesplits query all fee splits
	QueryFeesplits(pageReq *query.PageRequest) (*types.QueryFeeSplitsResponse, error)

	// QueryFeeSplit query a registered contract for fee distribution by hex address
	QueryFeeSplit(contractAddress string) (*types.QueryFeeSplitResponse, error)

	// QueryParams query the current feesplit module parameters
	QueryParams() (*types.QueryParamsResponse, error)

	// QueryDeployerFeeSplits query all contracts that a given deployer has registered for fee distribution
	QueryDeployerFeeSplits(deployerAddress string, pageReq *query.PageRequest) (*types.QueryDeployerFeeSplitsResponse, error)

	// QueryWithdrawerFeeSplits query all contracts that have been registered for fee distribution with a given withdrawer address
	QueryWithdrawerFeeSplits(withdrawerAddress string, pageReq *query.PageRequest) (*types.QueryWithdrawerFeeSplitsResponse, error)
}
