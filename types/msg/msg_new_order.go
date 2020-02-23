package msg

import (
	"encoding/json"
	"github.com/okex/okchain-go-sdk/types"
)

type MsgNewOrder struct {
	Sender   types.AccAddress
	Product  string
	Side     string
	Price    types.Dec
	Quantity types.Dec
}

func NewMsgNewOrder(sender types.AccAddress, product string, side string, price string, quantity string) MsgNewOrder {
	return MsgNewOrder{
		Sender:   sender,
		Product:  product,
		Side:     side,
		Price:    types.MustNewDecFromStr(price),
		Quantity: types.MustNewDecFromStr(quantity),
	}
}

func (msg MsgNewOrder) Route() string { return "" }

func (msg MsgNewOrder) Type() string { return "" }

func (msg MsgNewOrder) ValidateBasic() types.Error {
	return nil
}

func (msg MsgNewOrder) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return types.MustSortJSON(b)
}

func (msg MsgNewOrder) GetSigners() []types.AccAddress {
	return nil
}
