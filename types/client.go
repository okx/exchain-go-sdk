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
}

// ClientQuery shows the expected query behavior
type ClientQuery interface {
	rpc.SignClient
	Query(path string, key cmn.HexBytes) ([]byte, error)
	QueryStore(key cmn.HexBytes, storeName, endPath string) ([]byte, error)
	QuerySubspace(subspace []byte, storeName string) ([]cmn.KVPair, error)
}

// ClientQuery shows the expected tx behavior
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
}

// NewClientConfig creates a new instance of ClientConfig
func NewClientConfig(nodeURI string, broadcastMode BroadcastMode) ClientConfig {
	return ClientConfig{
		nodeURI,
		broadcastMode,
	}
}
