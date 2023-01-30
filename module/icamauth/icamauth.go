package icamauth

import (
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/icamauth/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/client/context"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	icamauth "github.com/okex/exchain/x/icamauth/types"
)

var _ gosdktypes.Module = (*icamauthClient)(nil)

type icamauthClient struct {
	gosdktypes.BaseClient
	protoCdc *codec.ProtoCodec
	context.CLIContext
}

// RegisterCodec registers the msg type in ammswap module
func (ac icamauthClient) RegisterCodec(cdc *codec.Codec) {
	icamauth.RegisterCodec(cdc)
}

// Name returns the module name
func (icamauthClient) Name() string {
	return types.ModuleName
}

// NewAmmSwapClient creates a new instance of ammswap client as implement
func NewIcamauthClient(baseClient gosdktypes.BaseClient, protoCdc *codec.ProtoCodec) exposed.Icamauth {
	clientCtx := context.NewCLIContext().WithNodeURI(baseClient.GetConfig().NodeURI)
	return icamauthClient{baseClient, protoCdc, clientCtx}
}
