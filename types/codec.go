package types

import (
	"github.com/tendermint/go-amino"
	cryptoamino "github.com/tendermint/tendermint/crypto/encoding/amino"
)

var MsgCdc = amino.NewCodec()

func init() {
	RegisterMsgCdc(MsgCdc)
	cryptoamino.RegisterAmino(MsgCdc)
	MsgCdc.Seal()
}

func RegisterMsgCdc(cdc *amino.Codec) {
	cdc.RegisterInterface((*Msg)(nil), nil)
	cdc.RegisterConcrete(MsgSend{}, "okchain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgNewOrders{}, "okchain/order/MsgNew", nil)
	cdc.RegisterConcrete(MsgCancelOrders{}, "okchain/order/MsgCancel", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "okchain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgMint{}, "okchain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgDelegate{}, "okchain/staking/MsgDelegate", nil)
	cdc.RegisterConcrete(MsgUndelegate{}, "okchain/staking/MsgUnDelegate", nil)
	cdc.RegisterConcrete(MsgVote{}, "okchain/staking/MsgVote", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "okchain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgUnjail{}, "cosmos-sdk/MsgUnjail", nil)
	cdc.RegisterConcrete(MsgCreateValidator{}, "okchain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "okchain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "okchain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "okchain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "okchain/staking/MsgUnbindProxy", nil)
	cdc.RegisterConcrete(MsgList{}, "okchain/dex/MsgList", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "okchain/dex/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "okchain/dex/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "okchain/dex/MsgTransferTradingPairOwnership", nil)

	cdc.RegisterConcrete(MsgTokenIssue{}, "okchain/token/MsgIssue", nil)

	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "cosmos-sdk/StdTx", nil)
}
