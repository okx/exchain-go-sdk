package types

// StdSignMsg is a convenience structure for passing along a Msg with the other requirements for a StdSignDoc before
// it is signed
type StdSignMsg struct {
	ChainID       string `json:"chain_id"`
	AccountNumber uint64 `json:"account_number"`
	Sequence      uint64 `json:"sequence"`
	Fee           StdFee `json:"fee"`
	Msgs          []Msg  `json:"msgs"`
	Memo          string `json:"memo"`
}

// Bytes gets message bytes
func (msg StdSignMsg) Bytes() []byte {
	return stdSignBytes(msg.ChainID, msg.AccountNumber, msg.Sequence, msg.Fee, msg.Msgs, msg.Memo)
}
