package exposed

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// Slashing shows the expected behavior for inner slashing client
type Slashing interface {
	sdk.Module
	SlashingTx
}

// SlashingTx shows the expected tx behavior for inner slashing client
type SlashingTx interface {
	Unjail(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}
