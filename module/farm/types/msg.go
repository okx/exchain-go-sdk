package types

import sdk "github.com/okex/okexchain-go-sdk/types"

// MsgCreatePool - structure for creating a farm pool
type MsgCreatePool struct {
	Owner        sdk.AccAddress `json:"owner"`
	PoolName     string         `json:"pool_name"`
	SymbolLocked string         `json:"locked_symbol"`
	YieldSymbol  string         `json:"yield_symbol"`
}

// NewMsgCreatePool creates a msg of create-pool
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

// NewMsgDestroyPool creates a msg of destroy-pool
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