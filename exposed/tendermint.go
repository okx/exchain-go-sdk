package exposed

import (
	"github.com/okex/okexchain-go-sdk/module/tendermint/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// Tendermint shows the expected behavior for inner tendermint client
type Tendermint interface {
	gosdktypes.Module
	TendermintQuery
}

// TendermintQuery shows the expected query behavior for inner tendermint client
type TendermintQuery interface {
	QueryBlock(height int64) (*types.Block, error)
	QueryBlockResults(height int64) (*types.ResultBlockResults, error)
	QueryCommitResult(height int64) (*types.ResultCommit, error)
	QueryValidatorsResult(height int64) (*types.ResultValidators, error)
	QueryTxResult(txHash []byte, prove bool) (types.ResultTx, error)
	// QueryTxsResult assumes the node to query a truth teller
	QueryTxsResult(queryStr string, page, perPage int) (types.ResultTxs, error)
}
