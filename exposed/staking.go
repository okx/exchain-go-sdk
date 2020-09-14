package exposed

import (
	"github.com/okex/okexchain-go-sdk/module/staking/types"
	sdk "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/types/crypto/keys"
)

// Staking shows the expected behavior for inner staking client
type Staking interface {
	sdk.Module
	StakingTx
	StakingQuery
}

// StakingTx shows the expected tx behavior for inner staking client
type StakingTx interface {
	CreateValidator(fromInfo keys.Info, passWd, pubkeyStr, moniker, identity, website, details, memo string, accNum,
		seqNum uint64) (sdk.TxResponse, error)
	DestroyValidator(fromInfo keys.Info, passWd string, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	EditValidator(fromInfo keys.Info, passWd, moniker, identity, website, details, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	Deposit(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	Withdraw(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	AddShares(fromInfo keys.Info, passWd string, valAddrsStr []string, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	RegisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	UnregisterProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	BindProxy(fromInfo keys.Info, passWd, proxyAddrStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
	UnbindProxy(fromInfo keys.Info, passWd, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// StakingQuery shows the expected query behavior for inner staking client
type StakingQuery interface {
	QueryValidators() ([]types.Validator, error)
	QueryValidator(valAddrStr string) (types.Validator, error)
	QueryDelegator(delAddrStr string) (types.DelegatorResp, error)
}
