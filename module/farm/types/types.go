package types

import sdk "github.com/okex/okexchain-go-sdk/types"

// const
const (
	ModuleName = "farm"

	QueryPoolPath             = "custom/farm/pool"
	QueryAccountPath          = "custom/farm/account"
	QueryAccountsLockedToPath = "custom/farm/accounts-locked-to"
	QueryLockInfoPath         = "custom/farm/lock-info"
)

var (
	msgCdc = sdk.NewCodec()
	// FarmPoolPrefix is useful for subspace and store query about pools
	FarmPoolPrefix = []byte{0x01}
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for farm module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgCreatePool{}, "okexchain/farm/MsgCreatePool")
	cdc.RegisterConcrete(MsgDestroyPool{}, "okexchain/farm/MsgDestroyPool")
	cdc.RegisterConcrete(MsgLock{}, "okexchain/farm/MsgLock")
	cdc.RegisterConcrete(MsgUnlock{}, "okexchain/farm/MsgUnlock")
	cdc.RegisterConcrete(MsgClaim{}, "okexchain/farm/MsgClaim")
	cdc.RegisterConcrete(MsgProvide{}, "okexchain/farm/MsgProvide")
	// MsgSetWhite is used for test
	// TODO: rm it later
	cdc.RegisterConcrete(MsgSetWhite{}, "okexchain/farm/MsgSetWhite")
}

// FarmPool is the pool where an address can lock specified token to yield other tokens
type FarmPool struct {
	Owner                   sdk.AccAddress    `json:"owner"`
	Name                    string            `json:"name"`
	MinLockAmount           sdk.DecCoin       `json:"min_lock_amount"`
	DepositAmount           sdk.DecCoin       `json:"deposit_amount"`
	TotalValueLocked        sdk.DecCoin       `json:"total_value_locked"`
	YieldedTokenInfos       YieldedTokenInfos `json:"yielded_token_infos"`
	TotalAccumulatedRewards sdk.DecCoins      `json:"total_accumulated_rewards"`
}

// YieldedTokenInfos is a collection of YieldedTokenInfo
type YieldedTokenInfos []YieldedTokenInfo

// YieldedTokenInfo is the token excluding native token which can be yielded by locking other tokens including LPT and
// token issued
type YieldedTokenInfo struct {
	RemainingAmount         sdk.DecCoin `json:"remaining_amount"`
	StartBlockHeightToYield int64       `json:"start_block_height_to_yield"`
	AmountYieldedPerBlock   sdk.Dec     `json:"amount_yielded_per_block"`
}

// LockInfo - structure of an address's locked info of
type LockInfo struct {
	Owner            sdk.AccAddress `json:"owner"`
	PoolName         string         `json:"pool_name"`
	Amount           sdk.DecCoin    `json:"amount"`
	StartBlockHeight int64          `json:"start_block_height"`
	ReferencePeriod  uint64         `json:"reference_period"`
}
