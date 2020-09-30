package exposed

import (
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/crypto/keys"
)

// Farm shows the expected behavior for inner farm client
type Farm interface {
	sdk.Module
	FarmTx
	FarmQuery
}

// FarmTx shows the expected tx behavior for inner farm client
type FarmTx interface {
	CreatePool(fromInfo keys.Info, passWd, poolName, lockToken, yieldToken, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	DestroyPool(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Provide(fromInfo keys.Info, passWd, poolName, amountStr, yieldPerBlockStr string, startHeightToYield int64, memo string, accNum,
		seqNum uint64) (sdk.TxResponse, error)
}

// FarmQuery shows the expected query behavior for inner farm client
type FarmQuery interface{}
