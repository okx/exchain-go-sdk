package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgNewOrders - structure for placing multi-orders
type MsgNewOrders struct {
	Sender     sdk.AccAddress `json:"sender"`
	OrderItems []OrderItem    `json:"order_items"`
}

// NewMsgNewOrders is a constructor function for MsgNewOrders
func NewMsgNewOrders(sender sdk.AccAddress, orderItems []OrderItem) MsgNewOrders {
	return MsgNewOrders{
		Sender:     sender,
		OrderItems: orderItems,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgNewOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgNewOrders) Route() string                { return "" }
func (MsgNewOrders) Type() string                 { return "" }
func (MsgNewOrders) ValidateBasic() sdk.Error     { return nil }
func (MsgNewOrders) GetSigners() []sdk.AccAddress { return nil }

// MsgCancelOrders - structure for canceling several orders that have been placed
type MsgCancelOrders struct {
	Sender   sdk.AccAddress `json:"sender"`
	OrderIDs []string       `json:"order_ids"`
}

// NewMsgCancelOrders is a constructor function for MsgCancelOrders
func NewMsgCancelOrders(sender sdk.AccAddress, orderIDItems []string) MsgCancelOrders {
	return MsgCancelOrders{
		Sender:   sender,
		OrderIDs: orderIDItems,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelOrders) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgCancelOrders) Route() string                { return "" }
func (MsgCancelOrders) Type() string                 { return "" }
func (MsgCancelOrders) ValidateBasic() sdk.Error     { return nil }
func (MsgCancelOrders) GetSigners() []sdk.AccAddress { return nil }
