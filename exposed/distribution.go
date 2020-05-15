package exposed

import (
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
)

// Distribution shows the expected behavior for inner distribution client
type Distribution interface {
	sdk.Module
	DistrTx
}

// DistrTx shows the expected tx behavior for inner distribution client
type DistrTx interface {
	SetWithdrawAddr(fromInfo keys.Info, passWd, withdrawAddrStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}