package types

import (
	"math/big"

	apptypes "github.com/okx/okbchain/app/types"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	evmtypes "github.com/okx/okbchain/x/evm/types"
)

// const
const (
	ModuleName      = evmtypes.ModuleName
	defaultGasPrice = "0.000000001"
)

type (
	QueryResCode    = evmtypes.QueryResCode
	QueryResStorage = evmtypes.QueryResStorage
)

var (
	DefaultGasPrice    = sdk.MustNewDecFromStr(defaultGasPrice).BigInt()
	DefaultRPCGasLimit = big.NewInt(apptypes.DefaultRPCGasLimit)
)
