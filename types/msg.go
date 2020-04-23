package types

// Msg shows the expected behavior of any msgs of OKChain
type Msg interface {
	Route() string
	Type() string
	ValidateBasic() Error
	GetSignBytes() []byte
	GetSigners() []AccAddress
}

// Tx shows the expected behavior of any txs of OKChain
type Tx interface {
	GetMsgs() []Msg
	ValidateBasic() Error
}

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, Error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)
