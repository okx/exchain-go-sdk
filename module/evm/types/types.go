package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	apptypes "github.com/okex/okexchain/app/types"
	evmtypes "github.com/okex/okexchain/x/evm/types"
	"math/big"
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
