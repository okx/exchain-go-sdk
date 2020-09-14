package exposed

import (
	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/crypto/keys"
)

// AmmSwap shows the expected behavior for inner staking client
type AmmSwap interface {
	sdk.Module
	AmmSwapTx
	AmmSwapQuery
}

// AmmSwapTx shows the expected tx behavior for inner staking client
type AmmSwapTx interface {
	AddLiquidity(fromInfo keys.Info, passWd, minLiquidity, maxBaseAmount, quoteAmount, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	RemoveLiquidity(fromInfo keys.Info, passWd, liquidity, minBaseAmount, minQuoteAmount, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	CreateExchange(fromInfo keys.Info, passWd, token, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	TokenSwap(fromInfo keys.Info, passWd, soldTokenAmount, minBoughtTokenAmount, recipient, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
}

// AmmSwapQuery shows the expected query behavior for inner staking client
type AmmSwapQuery interface {
	QueryExchange(token string) (types.SwapTokenPair, error)
}
