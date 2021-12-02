package types

import (
	"math/big"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	apptypes "github.com/okex/exchain/app/types"
	evmtypes "github.com/okex/exchain/x/evm/types"
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
