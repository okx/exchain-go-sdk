package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"github.com/okex/exchain/x/feesplit/types"
)

// Token shows the expected behavior for inner token client
type Feesplit interface {
	gosdktypes.Module
	FeesplitTx
	FeesplitQuery
}

// TokenTx shows the expected tx behavior for inner token client
type FeesplitTx interface {
}

// FeesplitQuery shows the expected query behavior for inner token client
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
