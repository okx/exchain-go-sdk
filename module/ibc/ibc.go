package ibc

import (
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"

	gosdktypes "github.com/okex/exchain-go-sdk/types"
)

var (
	_ gosdktypes.Module = (*ibcClient)(nil)
)

type ibcClient struct {
	gosdktypes.BaseClient
}

func (ibc ibcClient) RegisterCodec(cdc *codec.Codec) {
	//proto.RegisterType((*types.MsgTransfer)(nil), "/ibc.applications.transfer.v1.MsgTransfer")
	cdc.RegisterConcrete(types.MsgTransfer{}, "/ibc.applications.transfer.v1.MsgTransfer", nil)
}

// Name returns the module name
func (ibcClient) Name() string {
	return types.ModuleName
}

// NewIbcClient creates a new instance of auth client as implement
func NewIbcClient(baseClient gosdktypes.BaseClient) ibcClient {
	return ibcClient{baseClient}
}
