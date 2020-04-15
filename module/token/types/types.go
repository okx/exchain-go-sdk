package types

import "github.com/okex/okchain-go-sdk/types"

// const
const (
	ModuleName = "token"
)

var (
	msgCdc = types.NewCodec()
)

func init() {
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for token module
func RegisterCodec(cdc types.SDKCodec) {
	cdc.RegisterConcrete(MsgSend{}, "okchain/token/MsgTransfer")
	cdc.RegisterConcrete(MsgMultiSend{}, "okchain/token/MsgMultiTransfer")
	cdc.RegisterConcrete(MsgTokenIssue{}, "okchain/token/MsgIssue")
	cdc.RegisterConcrete(MsgMint{}, "okchain/token/MsgMint")
}

// TransferUnit - amount part for multi-send
type TransferUnit struct {
	To    types.AccAddress `json:"to"`
	Coins types.DecCoins   `json:"coins"`
}

// NewTransferUnit creates a new instance of TransferUnit
func NewTransferUnit(addr types.AccAddress, coins types.DecCoins) TransferUnit {
	return TransferUnit{
		To:    addr,
		Coins: coins,
	}
}
