package types

import sdk "github.com/okex/okexchain-go-sdk/types"

// MsgCreatePool - structure for creating a farm pool
type MsgCreatePool struct {
	Owner        sdk.AccAddress `json:"owner"`
	PoolName     string         `json:"pool_name"`
	SymbolLocked string         `json:"locked_symbol"`
	YieldSymbol  string         `json:"yield_symbol"`
}

// NewMsgCreatePool creates a new instance of MsgCreatePool
func NewMsgCreatePool(address sdk.AccAddress, poolName, lockToken, yieldToken string) MsgCreatePool {
	return MsgCreatePool{
		Owner:        address,
		PoolName:     poolName,
		SymbolLocked: lockToken,
		YieldSymbol:  yieldToken,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreatePool) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgCreatePool) Route() string                { return "" }
func (MsgCreatePool) Type() string                 { return "" }
func (MsgCreatePool) ValidateBasic() sdk.Error     { return nil }
func (MsgCreatePool) GetSigners() []sdk.AccAddress { return nil }

// MsgDestroyPool - structure for destroying a farm pool
type MsgDestroyPool struct {
	Owner    sdk.AccAddress `json:"owner"`
	PoolName string         `json:"pool_name"`
}

// NewMsgDestroyPool creates a new instance of MsgDestroyPool
func NewMsgDestroyPool(address sdk.AccAddress, poolName string) MsgDestroyPool {
	return MsgDestroyPool{
		Owner:    address,
		PoolName: poolName,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDestroyPool) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDestroyPool) Route() string                { return "" }
func (MsgDestroyPool) Type() string                 { return "" }
func (MsgDestroyPool) ValidateBasic() sdk.Error     { return nil }
func (MsgDestroyPool) GetSigners() []sdk.AccAddress { return nil }

// MsgProvide - structure for providing tokens as the yield into a farm pool
type MsgProvide struct {
	PoolName              string         `json:"pool_name"`
	Address               sdk.AccAddress `json:"address"`
	Amount                sdk.DecCoin    `json:"amount"`
	AmountYieldedPerBlock sdk.Dec        `json:"amount_yielded_per_block"`
	StartHeightToYield    int64          `json:"start_height_to_yield"`
}

// NewMsgProvide creates a new instance of MsgProvide
func NewMsgProvide(address sdk.AccAddress, poolName string, amount sdk.DecCoin, amountYieldedPerBlock sdk.Dec,
	startHeightToYield int64) MsgProvide {
	return MsgProvide{
		PoolName:              poolName,
		Address:               address,
		Amount:                amount,
		AmountYieldedPerBlock: amountYieldedPerBlock,
		StartHeightToYield:    startHeightToYield,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgProvide) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgProvide) Route() string                { return "" }
func (MsgProvide) Type() string                 { return "" }
func (MsgProvide) ValidateBasic() sdk.Error     { return nil }
func (MsgProvide) GetSigners() []sdk.AccAddress { return nil }

// MsgLock - structure for locking tokens into a farm pool
type MsgLock struct {
	PoolName string         `json:"pool_name"`
	Address  sdk.AccAddress `json:"address"`
	Amount   sdk.DecCoin    `json:"amount"`
}

// NewMsgLock creates a new instance of MsgLock
func NewMsgLock(address sdk.AccAddress, poolName string, amount sdk.DecCoin) MsgLock {
	return MsgLock{
		PoolName: poolName,
		Address:  address,
		Amount:   amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgLock) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgLock) Route() string                { return "" }
func (MsgLock) Type() string                 { return "" }
func (MsgLock) ValidateBasic() sdk.Error     { return nil }
func (MsgLock) GetSigners() []sdk.AccAddress { return nil }

// MsgUnlock - structure for unlocking tokens from a farm pool
type MsgUnlock struct {
	PoolName string         `json:"pool_name"`
	Address  sdk.AccAddress `json:"address"`
	Amount   sdk.DecCoin    `json:"amount"`
}

// NewMsgUnlock creates a new instance of MsgUnlock
func NewMsgUnlock(address sdk.AccAddress, poolName string, amount sdk.DecCoin) MsgUnlock {
	return MsgUnlock{
		PoolName: poolName,
		Address:  address,
		Amount:   amount,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgUnlock) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgUnlock) Route() string                { return "" }
func (MsgUnlock) Type() string                 { return "" }
func (MsgUnlock) ValidateBasic() sdk.Error     { return nil }
func (MsgUnlock) GetSigners() []sdk.AccAddress { return nil }

// MsgClaim - structure for claiming current yield
type MsgClaim struct {
	PoolName string         `json:"pool_name"`
	Address  sdk.AccAddress `json:"address"`
}

// NewMsgClaim creates a new instance of MsgClaim
func NewMsgClaim(address sdk.AccAddress, poolName string) MsgClaim {
	return MsgClaim{
		PoolName: poolName,
		Address:  address,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgClaim) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgClaim) Route() string                { return "" }
func (MsgClaim) Type() string                 { return "" }
func (MsgClaim) ValidateBasic() sdk.Error     { return nil }
func (MsgClaim) GetSigners() []sdk.AccAddress { return nil }
