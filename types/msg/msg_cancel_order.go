package msg

import (
	"encoding/json"
	"github.com/ok-chain/gosdk/types"
)

type MsgCancelOrder struct {
	Sender  types.AccAddress
	OrderId string
}

// NewMsgCancelOrder is a constructor function for MsgCancelOrder
func NewMsgCancelOrder(sender types.AccAddress, orderId string) MsgCancelOrder {
	msgCancelOrder := MsgCancelOrder{
		Sender:  sender,
		OrderId: orderId,
	}
	return msgCancelOrder
}

func (msg MsgCancelOrder) Route() string { return "" }

func (msg MsgCancelOrder) Type() string { return "" }

func (msg MsgCancelOrder) ValidateBasic() types.Error {
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelOrder) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return types.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgCancelOrder) GetSigners() []types.AccAddress {
	return nil
}
