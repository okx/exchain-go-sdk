package types

import (
	"errors"

	cmn "github.com/tendermint/tendermint/libs/common"
	rpc "github.com/tendermint/tendermint/rpc/client"
)

// BaseClient shows the expected behavior for a base client
type BaseClient interface {
	ClientQuery
	ClientTx
	TxHandler
	SimulationHandler
	GetCodec() SDKCodec
	GetConfig() ClientConfig
}

// TxHandler shows the expected behavior to handle tx
type TxHandler interface {
	BuildAndBroadcast(fromName, passphrase, memo string, msgs []Msg, accNumber, seqNumber uint64) (TxResponse, error)
	BuildStdTx(fromName, passphrase, memo string, msgs []Msg, accNumber, seqNumber uint64) (StdTx, error)
	BuildUnsignedStdTxOffline(msgs []Msg, memo string) StdTx
}

// SimulationHandler shows the expected behavior to handle simulation
type SimulationHandler interface {
	CalculateGas(txBytes []byte) (StdFee, error)
	BuildTxForSim(msgs []Msg, memo string, accNumber, seqNumber uint64) ([]byte, error)
}

// ClientQuery shows the expected query behavior
type ClientQuery interface {
	rpc.SignClient
	Query(path string, key cmn.HexBytes) ([]byte, error)
	QueryStore(key cmn.HexBytes, storeName, endPath string) ([]byte, error)
	QuerySubspace(subspace []byte, storeName string) ([]cmn.KVPair, error)
}

// ClientTx shows the expected tx behavior
type ClientTx interface {
	Broadcast(txBytes []byte, broadcastMode BroadcastMode) (res TxResponse, err error)
}

// RPCClient shows the expected behavior for a inner exposed client
type RPCClient interface {
	rpc.ABCIClient
	rpc.SignClient
}

// ClientConfig records the base config of gosdk client
type ClientConfig struct {
	NodeURI       string
	ChainID       string
	BroadcastMode BroadcastMode
	Gas           uint64
	GasAdjustment float64
	Fees          DecCoins
	GasPrices     DecCoins
}

// NewClientConfig creates a new instance of ClientConfig
func NewClientConfig(nodeURI, chainID string, broadcastMode BroadcastMode, feesStr string, gas uint64, gasAdjustment float64,
	gasPricesStr string) (
	cliConfig ClientConfig, err error) {
	var fees, gasPrices DecCoins
	if len(feesStr) != 0 {
		fees, err = ParseDecCoins(feesStr)
		if err != nil {
			return
		}
	}

	if len(gasPricesStr) != 0 {
		if gasAdjustment <= 1 {
			return cliConfig, errors.New("failed. gasAdjustment must be greater than 1 with the auto gas calculating")
		}

		gasPrices, err = ParseDecCoins(gasPricesStr)
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
