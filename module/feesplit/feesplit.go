package feesplit

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/client/context"
	"github.com/okex/exchain/libs/cosmos-sdk/codec"
	"github.com/okex/exchain/x/feesplit/types"
)

const (
	// Amino names
	registerFeeSplitName = "okexchain/MsgRegisterFeeSplit"
	updateFeeSplitName   = "okexchain/MsgUpdateFeeSplit"
	cancelFeeSplitName   = "okexchain/MsgCancelFeeSplit"
)

var (
	_ gosdktypes.Module = (*feesplitClient)(nil)
)

type feesplitClient struct {
	gosdktypes.BaseClient
	context.CLIContext
}

func (ibc feesplitClient) RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgRegisterFeeSplit{}, registerFeeSplitName, nil)
	cdc.RegisterConcrete(types.MsgUpdateFeeSplit{}, updateFeeSplitName, nil)
	cdc.RegisterConcrete(types.MsgCancelFeeSplit{}, cancelFeeSplitName, nil)
}

// Name returns the module name
func (feesplitClient) Name() string {
	return types.ModuleName
}

// NewfeesplitClient creates a new instance of auth client as implement
func NewfeesplitClient(baseClient gosdktypes.BaseClient) feesplitClient {
	clientCtx := context.NewCLIContext().WithNodeURI(baseClient.GetConfig().NodeURI)
	return feesplitClient{baseClient, clientCtx}
}
