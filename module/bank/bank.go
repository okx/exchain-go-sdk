package bank

import (
	"github.com/okex/exchain-go-sdk/module/ammswap/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain/libs/cosmos-sdk/x/bank"
)

var _ gosdktypes.Module = (*bankClient)(nil)

type bankClient struct {
	gosdktypes.BaseClient
}

// RegisterCodec registers the msg type in ammswap module
func (ac bankClient) RegisterCodec(cdc *codec.Codec) {
	bank.RegisterCodec(cdc)
	//proto.RegisterType((*types.MsgTransfer)(nil), "/ibc.applications.transfer.v1.MsgTransfer")
	//cdc.RegisterConcrete(types.MsgTransfer{}, "/ibc.applications.transfer.v1.MsgTransfer", nil)
}

// Name returns the module name
func (bankClient) Name() string {
	return types.ModuleName
}

//// NewAmmSwapClient creates a new instance of ammswap client as implement
//func NewAmmSwapClient(baseClient gosdktypes.BaseClient) exposed.AmmSwap {
//	//return bankClient{baseClient}
//}
