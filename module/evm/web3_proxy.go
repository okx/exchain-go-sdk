package evm

import (
	"errors"
	"fmt"
	"math/big"

	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/okex/exchain-go-sdk/exposed"
	"github.com/okex/exchain-go-sdk/module/evm/types"
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain-go-sdk/utils"
	rpctypes "github.com/okx/okbchain/app/rpc/types"
	apptypes "github.com/okx/okbchain/app/types"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/x/auth"
	"github.com/okx/okbchain/libs/cosmos-sdk/x/auth/exported"
	authtypes "github.com/okx/okbchain/libs/cosmos-sdk/x/auth/types"
	evmtypes "github.com/okx/okbchain/x/evm/types"
)

// Web3Proxy returns the client with exposed.Web3Proxy's behaviour
func (ec evmClient) Web3Proxy() exposed.Web3Proxy {
	return gosdktypes.Module(ec).(exposed.Web3Proxy)
}

// BlockNumberProxy returns the current block number as method "eth_blockNumber" without rest server routing
func (ec evmClient) BlockNumberProxy() (hexutil.Uint64, error) {
	resBlockchainInfo, err := ec.BlockchainInfo(0, 0)
	if err != nil {
		return hexutil.Uint64(0), err
	}

	blockNumber := resBlockchainInfo.LastHeight
	if blockNumber > 0 {
		// decrease blockNumber to make sure every block has been executed in local
		blockNumber--
	}

	return hexutil.Uint64(blockNumber), nil
}

// EstimateGasProxy returns the estimated gas according to the args as method "eth_estimateGas" without rest server routing
func (ec evmClient) EstimateGasProxy(args rpctypes.CallArgs) (estimatedGas hexutil.Uint64, err error) {
	simResponse, err := ec.doCallProxy(args, types.DefaultRPCGasLimit)
	if err != nil {
		return
	}

	estimatedGas = hexutil.Uint64(simResponse.GasInfo.GasUsed + 1000)
	return
}

func (ec evmClient) accountNonce(addr ethcmn.Address) (nonce uint64, err error) {
	// get account info from chain
	path := fmt.Sprintf("custom/%s/%s", auth.QuerierRoute, auth.QueryAccount)
	bytes, err := ec.GetCodec().MarshalJSON(authtypes.NewQueryAccountParams(addr.Bytes()))
	if err != nil {
		return nonce, utils.ErrClientQuery(err.Error())
	}

	res, _, err := ec.Query(path, bytes)
	if res == nil {
		return nonce, errors.New("failed. your account has no record on the chain")
	}

	var account exported.Account
	if err = ec.GetCodec().UnmarshalJSON(res, &account); err != nil {
		return nonce, utils.ErrUnmarshalJSON(err.Error())
	}

	return account.GetSequence(), err
}

func (ec evmClient) doCallProxy(args rpctypes.CallArgs, globalGasCap *big.Int) (
	simResponse *sdk.SimulationResponse, err error) {
	if args.From == nil {
		return simResponse, errors.New("failed. empty from address in CallArgs")
	}

	nonce, _ := ec.accountNonce(*args.From)

	// Set default gas & gas price if none were set
	// Change this to uint64(math.MaxUint64 / 2) if gas cap can be configured
	gas := uint64(apptypes.DefaultRPCGasLimit)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}

	if globalGasCap != nil && globalGasCap.Uint64() < gas {
		gas = globalGasCap.Uint64()
	}

	// Set gas price using default or parameter if passed in
	gasPrice := big.NewInt(apptypes.DefaultGasPrice)
	if args.GasPrice != nil {
		gasPrice = args.GasPrice.ToInt()
	}

	// Set value for transaction
	value := new(big.Int)
	if args.Value != nil {
		value = args.Value.ToInt()
	}

	// Set Data if provided
	var data []byte
	if args.Data != nil {
		data = *args.Data
	}

	simMsg := evmtypes.NewMsgEthereumTx(nonce, args.To, value, gas, gasPrice, data)

	var stdSig authtypes.StdSignature
	tx := authtypes.NewStdTx([]sdk.Msg{simMsg}, authtypes.StdFee{}, []authtypes.StdSignature{stdSig}, "")
	if err = tx.ValidateBasic(); err != nil {
		return
	}

	cdc := ec.GetCodec()
	txBytes, err := cdc.MarshalBinaryLengthPrefixed(tx)
	if err != nil {
		return
	}

	// Transaction simulation through query
	res, _, err := ec.Query("app/simulate", txBytes)
	if err != nil {
		return
	}

	simResponse = new(sdk.SimulationResponse)

	return simResponse, cdc.UnmarshalBinaryBare(res, simResponse)
}
