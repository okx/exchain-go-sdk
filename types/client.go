package types

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

// BaseClient shows the expected behavior for a base client
type BaseClient interface {
	ClientQuery
	ClientTx
	TxHandler
	SimulationHandler
	GetCodec() *codec.Codec
	GetConfig() ClientConfig
}

// TxHandler shows the expected behavior to handle tx
type TxHandler interface {
	BuildAndBroadcast(fromName, passphrase, memo string, msgs []sdk.Msg, accNumber, seqNumber uint64) (sdk.TxResponse, error)
	BuildStdTx(fromName, passphrase, memo string, msgs []sdk.Msg, accNumber, seqNumber uint64) (authtypes.StdTx, error)
	BuildUnsignedStdTxOffline(msgs []sdk.Msg, memo string) authtypes.StdTx
}

// SimulationHandler shows the expected behavior to handle simulation
type SimulationHandler interface {
	CalculateGas(txBytes []byte) (authtypes.StdFee, error)
	BuildTxForSim(msgs []sdk.Msg, memo string, accNumber, seqNumber uint64) ([]byte, error)
}

// ClientQuery shows the expected query behavior
type ClientQuery interface {
	rpcclient.SignClient
	Query(path string, key tmbytes.HexBytes) ([]byte, int64, error)
	QueryStore(key tmbytes.HexBytes, storeName, endPath string) ([]byte, int64, error)
}

// ClientTx shows the expected tx behavior
type ClientTx interface {
	Broadcast(txBytes []byte, broadcastMode string) (res sdk.TxResponse, err error)
}

// RPCClient shows the expected behavior for a inner exposed client
type RPCClient interface {
	rpcclient.ABCIClient
	rpcclient.SignClient
}

// ClientConfig records the base config of gosdk client
type ClientConfig struct {
	NodeURI       string
	ChainID       string
	BroadcastMode string
	Gas           uint64
	GasAdjustment float64
	Fees          sdk.DecCoins
	GasPrices     sdk.DecCoins
}

// NewClientConfig creates a new instance of ClientConfig
func NewClientConfig(nodeURI, chainID string, broadcastMode string, feesStr string, gas uint64, gasAdjustment float64,
	gasPricesStr string) (cliConfig ClientConfig, err error) {
	var fees, gasPrices sdk.DecCoins
	if len(feesStr) != 0 {
		fees, err = sdk.ParseDecCoins(feesStr)
		if err != nil {
			return
		}
	}

	if len(gasPricesStr) != 0 {
		if gasAdjustment <= 1 {
			return cliConfig, errors.New("failed. gasAdjustment must be greater than 1 with the auto gas calculating")
		}

		gasPrices, err = sdk.ParseDecCoins(gasPricesStr)
		if err != nil {
			return
		}
	}

	return ClientConfig{
		NodeURI:       nodeURI,
		ChainID:       chainID,
		BroadcastMode: broadcastMode,
		Gas:           gas,
		GasAdjustment: gasAdjustment,
		Fees:          fees,
		GasPrices:     gasPrices,
	}, err
}
