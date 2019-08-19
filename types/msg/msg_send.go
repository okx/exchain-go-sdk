package msg

import (
	"encoding/json"
	"github.com/ok-chain/ok-gosdk/types"
)


const (
	RouterKey = "token"
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
func (msg MsgSend) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSend) Type() string { return "send" }

// ValidateBasic Implements Msg.
func (msg MsgSend) ValidateBasic() types.Error {
	if msg.FromAddress.Empty() {
		return types.ErrInvalidAddress("missing sender address")
	}
	if msg.ToAddress.Empty() {
		return types.ErrInvalidAddress("missing recipient address")
	}
	if !msg.Amount.IsValid() {
		return types.ErrInvalidCoins("send amount is invalid: " + msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return types.ErrInsufficientCoins("send amount must be positive")
	}
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
	return []types.AccAddress{msg.FromAddress}
}
