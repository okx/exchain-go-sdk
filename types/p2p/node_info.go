package p2p

import (
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/version"
	"time"
)

type DefaultNodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`

	// Authenticate
	// TODO: replace with NetAddress
	ID_        ID     `json:"id"`          // authenticated identifier
	ListenAddr string `json:"listen_addr"` // accepting incoming

	// Check compatibility.
	// Channels are HexBytes so easier to read as JSON
	Network  string       `json:"network"`  // network/chain ID
	Version  string       `json:"version"`  // major.minor.revision
	Channels cmn.HexBytes `json:"channels"` // channels this node knows about

	// ASCIIText fields
	Moniker string               `json:"moniker"` // arbitrary moniker
	Other   DefaultNodeInfoOther `json:"other"`   // other application specific data
}

type ProtocolVersion struct {
	P2P   version.Protocol `json:"p2p"`
	Block version.Protocol `json:"block"`
	App   version.Protocol `json:"app"`
}

type ID string

type DefaultNodeInfoOther struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

type ConnectionStatus struct {
	Duration    time.Duration
	SendMonitor Status
	RecvMonitor Status
	Channels    []ChannelStatus
}

type Status struct {
	Active   bool          // Flag indicating an active transfer
	Start    time.Time     // Transfer start time
	Duration time.Duration // Time period covered by the statistics
	Idle     time.Duration // Time since the last transfer of at least 1 byte
	Bytes    int64         // Total number of bytes transferred
	Samples  int64         // Total number of samples taken
	InstRate int64         // Instantaneous transfer rate
	CurRate  int64         // Current transfer rate (EMA of InstRate)
	AvgRate  int64         // Average transfer rate (Bytes / Duration)
	PeakRate int64         // Maximum instantaneous transfer rate
	BytesRem int64         // Number of bytes remaining in the transfer
	TimeRem  time.Duration // Estimated time to completion
	Progress Percent       // Overall transfer progress
}

type Percent uint32

type ChannelStatus struct {
	ID                byte
	SendQueueCapacity int
	SendQueueSize     int
	Priority          int
	RecentlySent      int64
}
