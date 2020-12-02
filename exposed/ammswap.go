package exposed

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain-go-sdk/module/ammswap/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
)

// AmmSwap shows the expected behavior for inner ammswap client
type AmmSwap interface {
	gosdktypes.Module
	AmmSwapTx
	AmmSwapQuery
}

// AmmSwapTx shows the expected tx behavior for inner ammswap client
type AmmSwapTx interface {
	AddLiquidity(fromInfo keys.Info, passWd, minLiquidity, maxBaseAmount, quoteAmount, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	RemoveLiquidity(fromInfo keys.Info, passWd, liquidity, minBaseAmount, minQuoteAmount, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	CreateExchange(fromInfo keys.Info, passWd, baseToken, quoteToken, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	TokenSwap(fromInfo keys.Info, passWd, soldTokenAmount, minBoughtTokenAmount, recipient, deadlineDuration, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
}

// AmmSwapQuery shows the expected query behavior for inner ammswap client
type AmmSwapQuery interface {
	QuerySwapTokenPair(token string) (types.SwapTokenPair, error)
	QuerySwapTokenPairs() ([]types.SwapTokenPair, error)
	QueryBuyAmount(soldToken sdk.DecCoin, tokenToBuy string) (sdk.Dec, error)
}
