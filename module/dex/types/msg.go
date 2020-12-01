package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// MsgWithdraw - structure for withdrawing from a product
type MsgWithdraw struct {
	Product   string         `json:"product"`
	Amount    sdk.DecCoin    `json:"amount"`
	Depositor sdk.AccAddress `json:"depositor"`
}

// NewMsgWithdraw creates a msg of withdrawing
func NewMsgWithdraw(depositor sdk.AccAddress, product string, amount sdk.DecCoin) MsgWithdraw {
	return MsgWithdraw{
		Product:   product,
		Amount:    amount,
		Depositor: depositor,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgWithdraw) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgWithdraw) Route() string                { return "" }
func (MsgWithdraw) Type() string                 { return "" }
func (MsgWithdraw) ValidateBasic() sdk.Error     { return nil }
func (MsgWithdraw) GetSigners() []sdk.AccAddress { return nil }

// MsgTransferOwnership - structure to change the owner of the product
type MsgTransferOwnership struct {
	FromAddress sdk.AccAddress   `json:"from_address"`
	ToAddress   sdk.AccAddress   `json:"to_address"`
	Product     string           `json:"product"`
	ToSignature authtypes.StdSignature `json:"to_signature"`
}

// NewMsgTransferOwnership creates a msg of changing product's owner
func NewMsgTransferOwnership(fromAddr, toAddr sdk.AccAddress, product string) MsgTransferOwnership {
	return MsgTransferOwnership{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Product:     product,
		ToSignature: authtypes.StdSignature{},
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgTransferOwnership) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgTransferOwnership) Route() string                { return "" }
func (MsgTransferOwnership) Type() string                 { return "" }
func (MsgTransferOwnership) ValidateBasic() sdk.Error     { return nil }
func (MsgTransferOwnership) GetSigners() []sdk.AccAddress { return nil }
