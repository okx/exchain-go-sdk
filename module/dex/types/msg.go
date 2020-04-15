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
