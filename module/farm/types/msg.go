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
