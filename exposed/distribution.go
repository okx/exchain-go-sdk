package exposed

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
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
