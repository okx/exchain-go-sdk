package exposed

import (
	"github.com/okex/okexchain-go-sdk/module/farm/types"
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
	Lock(fromInfo keys.Info, passWd, poolName, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Unlock(fromInfo keys.Info, passWd, poolName, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Claim(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	// SetWhite is used for test
	// TODO: rm it later
	SetWhite(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// FarmQuery shows the expected query behavior for inner farm client
type FarmQuery interface {
	QueryPools() ([]types.FarmPool, error)
	QueryPool(poolName string) (types.FarmPool, error)
}
