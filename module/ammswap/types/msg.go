package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// MsgTokenToNativeToken defines the message for swap between token and DefaultBondDenom
type MsgTokenToNativeToken struct {
	SoldTokenAmount      sdk.DecCoin    `json:"sold_token_amount"`       // Amount of Tokens sold.
	MinBoughtTokenAmount sdk.DecCoin    `json:"min_bought_token_amount"` // Minimum token purchased.
	Deadline             int64          `json:"deadline"`                // Time after which this transaction can no longer be executed.
	Recipient            sdk.AccAddress `json:"recipient"`               // Recipient address,transfer Tokens to recipient.default recipient is sender.
	Sender               sdk.AccAddress `json:"sender"`                  // Sender
}

// NewMsgTokenToNativeToken is a constructor function for MsgTokenOKTSwap
func NewMsgTokenToNativeToken(
	soldTokenAmount, minBoughtTokenAmount sdk.DecCoin, deadline int64, recipient, sender sdk.AccAddress,
) MsgTokenToNativeToken {
	return MsgTokenToNativeToken{
		SoldTokenAmount:      soldTokenAmount,
		MinBoughtTokenAmount: minBoughtTokenAmount,
		Deadline:             deadline,
		Recipient:            recipient,
		Sender:               sender,
	}
}

// Route returns the name of the module
func (msg MsgTokenToNativeToken) Route() string { return "" }

// Type returns the action
func (msg MsgTokenToNativeToken) Type() string { return "" }

// ValidateBasic runs stateless checks on the message
func (msg MsgTokenToNativeToken) ValidateBasic() sdk.Error { return nil }

// GetSignBytes encodes the message for signing
func (msg MsgTokenToNativeToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgTokenToNativeToken) GetSigners() []sdk.AccAddress { return nil }

// GetSwapTokenPair defines token pair
func (msg MsgTokenToNativeToken) GetSwapTokenPair() string {
	if msg.SoldTokenAmount.Denom == "okt" {
		return msg.MinBoughtTokenAmount.Denom + "_" + msg.SoldTokenAmount.Denom
	}
	return msg.SoldTokenAmount.Denom + "_" + msg.MinBoughtTokenAmount.Denom
}
