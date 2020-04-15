package exposed

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// Dex shows the expected behavior for inner dex client
type Dex interface {
	sdk.Module
	DexTx
	DexQuery
}

// DexTx shows the expected tx behavior for inner dex client
type DexTx interface {
	List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	Deposit(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Withdraw(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// DexQuery shows the expected query behavior for inner dex client
type DexQuery interface {
}
