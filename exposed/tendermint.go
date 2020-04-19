package exposed

import (
	"github.com/okex/okchain-go-sdk/module/tendermint/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// Tendermint shows the expected behavior for inner tendermint client
type Tendermint interface {
	sdk.Module
	TendermintQuery
}

// TendermintQuery shows the expected query behavior for inner tendermint client
type TendermintQuery interface {
	QueryBlock(height int64) (types.Block, error)
}
