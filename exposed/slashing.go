package exposed

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	"github.com/okex/okchain-go-sdk/types"
)

// Slashing shows the expected behavior for inner slashing client
type Slashing interface {
	types.Module
	SlashingTx
}

// SlashingTx shows the expected tx behavior for inner slashing client
type SlashingTx interface {
	Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (types.TxResponse, error)
}
