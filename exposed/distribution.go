package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
)

// Distribution shows the expected behavior for inner distribution client
type Distribution interface {
	gosdktypes.Module
	DistrTx
}

// DistrTx shows the expected tx behavior for inner distribution client
type DistrTx interface {
	SetWithdrawAddr(fromInfo keys.Info, passWd, withdrawAddrStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	WithdrawRewards(fromInfo keys.Info, passWd, valAddrStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}
