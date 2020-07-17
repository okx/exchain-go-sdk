package exposed

import (
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
)

// PoolSwap shows the expected behavior for inner staking client
type PoolSwap interface {
	sdk.Module
	PoolSwapTx
	PoolSwapQuery
}

// PoolSwapTx shows the expected tx behavior for inner staking client
type PoolSwapTx interface {
	AddLiquidity(fromInfo keys.Info, passWd, minLiquidity, maxBaseAmount, quoteAmount, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	RemoveLiquidity(fromInfo keys.Info, passWd, liquidity, minBaseAmount, minQuoteAmount, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	CreateExchange(fromInfo keys.Info, passWd, token, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	TokenSwap(fromInfo keys.Info, passWd, soldTokenAmount, minBoughtTokenAmount, recipient, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
}

// PoolSwapQuery shows the expected query behavior for inner staking client
type PoolSwapQuery interface {

}
