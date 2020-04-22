package types

import (
	cmn "github.com/tendermint/tendermint/libs/common"
	rpc "github.com/tendermint/tendermint/rpc/client"
)

// BaseClient shows the expected behavior for a base client
type BaseClient interface {
	ClientQuery
	ClientTx
	TxHandler
	GetCodec() SDKCodec
	GetConfig() ClientConfig
}

// TxHandler shows the expected behavior to handle tx
type TxHandler interface {
	BuildAndBroadcast(fromName, passphrase, memo string, msgs []Msg, accNumber, seqNumber uint64) (TxResponse, error)
	BuildStdTx(fromName, passphrase, memo string, msgs []Msg, accNumber, seqNumber uint64) (StdTx, error)
	BuildUnsignedStdTxOffline(msgs []Msg, memo string) StdTx
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
	BroadcastMode BroadcastMode
	ChainID       string
	Fees          DecCoins
	Gas           uint64
}

// NewClientConfig creates a new instance of ClientConfig
func NewClientConfig(nodeURI, chainID string, broadcastMode BroadcastMode, feesStr string, gas uint64) (
	cliConfig ClientConfig, err error) {
	fees, err := ParseDecCoins(feesStr)
	if err != nil {
		return
	}

	return ClientConfig{
		NodeURI:       nodeURI,
		BroadcastMode: broadcastMode,
		ChainID:       chainID,
		Fees:          fees,
		Gas:           gas,
	}, err
}
