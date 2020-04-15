package types

// Transactions messages must fulfill the Msg
type Msg interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Route() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}

//__________________________________________________________

// Transactions objects must fulfill the Tx
type Tx interface {
	// Gets the all the transaction's messages.
	GetMsgs() []Msg

	// ValidateBasic does a simple and lightweight validation check that doesn't
	// require access to any other information.
	ValidateBasic() Error
}

//__________________________________________________________

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, Error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)

//__________________________________________________________
// msgs of okchain's tx

type MsgNewOrders struct {
	Sender     AccAddress  `json:"sender"`
	OrderItems []OrderItem `json:"order_items"`
}

// NewMsgNewOrders is a constructor function for MsgNewOrders
func NewMsgNewOrders(sender AccAddress, orderItems []OrderItem) MsgNewOrders {
	return MsgNewOrders{
		Sender:     sender,
		OrderItems: orderItems,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgNewOrders) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgNewOrders) Route() string            { return "" }
func (MsgNewOrders) Type() string             { return "" }
func (MsgNewOrders) ValidateBasic() Error     { return nil }
func (MsgNewOrders) GetSigners() []AccAddress { return nil }

type MsgCancelOrders struct {
	Sender   AccAddress `json:"sender"`
	OrderIds []string   `json:"order_ids"`
}

// NewMsgCancelOrders is a constructor function for MsgCancelOrders
func NewMsgCancelOrders(sender AccAddress, orderIdItems []string) MsgCancelOrders {
	return MsgCancelOrders{
		Sender:   sender,
		OrderIds: orderIdItems,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelOrders) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgCancelOrders) Route() string            { return "" }
func (MsgCancelOrders) Type() string             { return "" }
func (MsgCancelOrders) ValidateBasic() Error     { return nil }
func (MsgCancelOrders) GetSigners() []AccAddress { return nil }


// MsgList is a msg struct to list a trading pair on dex
type MsgList struct {
	Owner      AccAddress `json:"owner"`
	ListAsset  string     `json:"list_asset"`
	QuoteAsset string     `json:"quote_asset"`
	InitPrice  Dec        `json:"init_price"`
}

// NewMsgList creates a msg of listing a trading pair on dex
func NewMsgList(owner AccAddress, listAsset, quoteAsset string, initPrice Dec) MsgList {
	return MsgList{
		Owner:      owner,
		ListAsset:  listAsset,
		QuoteAsset: quoteAsset,
		InitPrice:  initPrice,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgList) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgList) Route() string            { return "" }
func (MsgList) Type() string             { return "" }
func (MsgList) ValidateBasic() Error     { return nil }
func (MsgList) GetSigners() []AccAddress { return nil }

// MsgList is a msg struct to deposit on a product
type MsgDeposit struct {
	Product   string     `json:"product"`
	Amount    DecCoin    `json:"amount"`
	Depositor AccAddress `json:"depositor"`
}

// NewMsgDeposit creates a msg of depositing
func NewMsgDeposit(depositor AccAddress, product string, amount DecCoin) MsgDeposit {
	return MsgDeposit{
		Product:   product,
		Amount:    amount,
		Depositor: depositor,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgDeposit) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgDeposit) Route() string            { return "" }
func (MsgDeposit) Type() string             { return "" }
func (MsgDeposit) ValidateBasic() Error     { return nil }
func (MsgDeposit) GetSigners() []AccAddress { return nil }

// MsgWithdraw is a msg struct to withdraw from a product
type MsgWithdraw struct {
	Product   string     `json:"product"`
	Amount    DecCoin    `json:"amount"`
	Depositor AccAddress `json:"depositor"`
}

// NewMsgWithdraw creates a msg of withdrawing
func NewMsgWithdraw(depositor AccAddress, product string, amount DecCoin) MsgWithdraw {
	return MsgWithdraw{
		Product:   product,
		Amount:    amount,
		Depositor: depositor,
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgWithdraw) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgWithdraw) Route() string            { return "" }
func (MsgWithdraw) Type() string             { return "" }
func (MsgWithdraw) ValidateBasic() Error     { return nil }
func (MsgWithdraw) GetSigners() []AccAddress { return nil }

// MsgTransferOwnership is a msg struct to change the owner of the product
type MsgTransferOwnership struct {
	FromAddress AccAddress   `json:"from_address"`
	ToAddress   AccAddress   `json:"to_address"`
	Product     string       `json:"product"`
	ToSignature StdSignature `json:"to_signature"`
}

// NewMsgTransferOwnership creates a msg of changing product's owner
func NewMsgTransferOwnership(from, to AccAddress, product string) MsgTransferOwnership {
	return MsgTransferOwnership{
		FromAddress: from,
		ToAddress:   to,
		Product:     product,
		ToSignature: StdSignature{},
	}
}

// GetSignBytes encodes the message for signing
func (msg MsgTransferOwnership) GetSignBytes() []byte {
	return MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// nolint
func (MsgTransferOwnership) Route() string            { return "" }
func (MsgTransferOwnership) Type() string             { return "" }
func (MsgTransferOwnership) ValidateBasic() Error     { return nil }
func (MsgTransferOwnership) GetSigners() []AccAddress { return nil }
