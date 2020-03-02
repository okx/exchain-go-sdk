package msg

import (
	"encoding/json"
	"github.com/okex/okchain-go-sdk/types"
)

type MsgMultiSend struct {
	From      types.AccAddress
	Transfers []types.TransferUnit
}

func NewMsgMultiSend(from types.AccAddress, transfers []types.TransferUnit) MsgMultiSend {
	return MsgMultiSend{
		From:      from,
		Transfers: transfers,
	}
}

func (msg MsgMultiSend) Route() string { return "" }

func (msg MsgMultiSend) Type() string { return "" }

func (msg MsgMultiSend) ValidateBasic() types.Error {
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgMultiSend) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return types.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgMultiSend) GetSigners() []types.AccAddress {
	return nil
}
