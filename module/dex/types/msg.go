package types

import (
	sdk "github.com/okex/okchain-go-sdk/types"
)

// MsgList - structure for listing a trading pair on dex
type MsgList struct {
	Owner      sdk.AccAddress `json:"owner"`
	ListAsset  string         `json:"list_asset"`
	QuoteAsset string         `json:"quote_asset"`
	InitPrice  sdk.Dec        `json:"init_price"`
}

// NewMsgList creates a msg of listing a trading pair on dex
func NewMsgList(owner sdk.AccAddress, listAsset, quoteAsset string, initPrice sdk.Dec) MsgList {
	return MsgList{
		Owner:      owner,
		ListAsset:  listAsset,
		QuoteAsset: quoteAsset,
		InitPrice:  initPrice,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgList) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgList) Route() string                { return "" }
func (MsgList) Type() string                 { return "" }
func (MsgList) ValidateBasic() sdk.Error     { return nil }
func (MsgList) GetSigners() []sdk.AccAddress { return nil }

// MsgList - structure for depositing on a product
type MsgDeposit struct {
	Product   string         `json:"product"`
	Amount    sdk.DecCoin    `json:"amount"`
	Depositor sdk.AccAddress `json:"depositor"`
}

// NewMsgDeposit creates a msg of depositing
func NewMsgDeposit(depositor sdk.AccAddress, product string, amount sdk.DecCoin) MsgDeposit {
	return MsgDeposit{
		Product:   product,
		Amount:    amount,
		Depositor: depositor,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDeposit) GetSignBytes() []byte {
	return sdk.MustSortJSON(msgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDeposit) Route() string                { return "" }
func (MsgDeposit) Type() string                 { return "" }
func (MsgDeposit) ValidateBasic() sdk.Error     { return nil }
func (MsgDeposit) GetSigners() []sdk.AccAddress { return nil }

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
	ToSignature sdk.StdSignature `json:"to_signature"`
}

// NewMsgTransferOwnership creates a msg of changing product's owner
func NewMsgTransferOwnership(fromAddr, toAddr sdk.AccAddress, product string) MsgTransferOwnership {
	return MsgTransferOwnership{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		Product:     product,
		ToSignature: sdk.StdSignature{},
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
