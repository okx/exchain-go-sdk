package exposed

import (
	"github.com/okex/exchain/libs/cosmos-sdk/crypto/keys"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain-go-sdk/module/farm/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
)

// Farm shows the expected behavior for inner farm client
type Farm interface {
	gosdktypes.Module
	FarmTx
	FarmQuery
}

// FarmTx shows the expected tx behavior for inner farm client
type FarmTx interface {
	CreatePool(fromInfo keys.Info, passWd, poolName, minLockAmountStr, yieldToken, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	DestroyPool(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Provide(fromInfo keys.Info, passWd, poolName, amountStr, yieldPerBlockStr string, startHeightToYield int64, memo string,
		accNum, seqNum uint64) (sdk.TxResponse, error)
	Lock(fromInfo keys.Info, passWd, poolName, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Unlock(fromInfo keys.Info, passWd, poolName, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Claim(fromInfo keys.Info, passWd, poolName, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// FarmQuery shows the expected query behavior for inner farm client
type FarmQuery interface {
	QueryPools() ([]types.FarmPool, error)
	QueryPool(poolName string) (types.FarmPool, error)
	QueryAccount(accAddrStr string) ([]string, error)
	QueryAccountsLockedTo(poolName string) ([]sdk.AccAddress, error)
	QueryLockInfo(poolName, accAddrStr string) (types.LockInfo, error)
	QueryEarnings(poolName, accAddrStr string) (types.Earnings, error)
}
