package exposed

import (
	"github.com/okex/okchain-go-sdk/crypto/keys"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// Token shows the expected behavior for inner token client
type Token interface {
	sdk.Module
	TokenTx
}

// TokenTx shows the expected tx behavior for inner token client
type TokenTx interface {
	Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}
