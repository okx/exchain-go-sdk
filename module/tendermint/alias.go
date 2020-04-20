package tendermint

import "github.com/okex/okchain-go-sdk/module/tendermint/types"

// const
const (
	ModuleName = types.ModuleName
)

type (
	// nolint
	Block = types.Block
	BlockResults = types.BlockResults
	ResultCommit = types.ResultCommit
	ResultValidators = types.ResultValidators
	ResultTx = types.ResultTx
	ResultTxs = types.ResultTxs
)
