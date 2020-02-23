package msg

import (
	"encoding/json"
	"github.com/okex/okchain-go-sdk/types"
)

type MsgSend struct {
	FromAddress types.AccAddress `json:"from_address"`
	ToAddress   types.AccAddress `json:"to_address"`
	Amount      types.Coins      `json:"amount"`
}

func NewMsgTokenSend(from, to types.AccAddress, coins types.Coins) MsgSend {
	return MsgSend{
		FromAddress: from,
		ToAddress:   to,
		Amount:      coins,
	}
}

// Route Implements Msg.
func (msg MsgSend) Route() string { return "" }

// Type Implements Msg.
func (msg MsgSend) Type() string { return "" }

// ValidateBasic Implements Msg.
func (msg MsgSend) ValidateBasic() types.Error {
	return nil
}

func (msg MsgSend) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return types.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSend) GetSigners() []types.AccAddress {
	return nil
}
