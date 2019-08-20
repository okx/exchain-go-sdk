package msg

import (
	"encoding/json"
	"github.com/ok-chain/gosdk/types"
)

type MsgMint struct {
	Symbol string
	Amount int64
	Owner  types.AccAddress
}

func NewMsgMint(symbol string, amount int64, owner types.AccAddress) MsgMint {
	return MsgMint{
		Symbol: symbol,
		Amount: amount,
		Owner:  owner,
	}
}

func (msg MsgMint) Route() string { return "" }

func (msg MsgMint) Type() string { return "" }

func (msg MsgMint) ValidateBasic() types.Error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMint) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return types.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgMint) GetSigners() []types.AccAddress {
	return nil
}
