package types

// Result is the union of ResponseFormat and ResponseCheckTx.
type Result struct {
	// Code is the response code, is stored back on the chain.
	Code CodeType

	// Codespace is the string referring to the domain of an error
	Codespace CodespaceType

	// Data is any data returned from the app.
	// Data has to be length prefixed in order to separate
	// results from multiple msgs executions
	Data []byte

	// Log contains the txs log information. NOTE: nondeterministic.
	Log string

	// GasWanted is the maximum units of work we allow this tx to perform.
	GasWanted uint64

	// GasUsed is the amount of gas actually consumed. NOTE: unimplemented
	GasUsed uint64

	// Tags are used for transaction indexing and pubsub.
	Tags Tags
}

// TODO: In the future, more codes may be OK.
func (res Result) IsOK() bool {
	return res.Code.IsOK()
}

// TxResponse defines a structure containing relevant tx data and metadata. The
// tags are stringified and the log is JSON decoded.
type TxResponse struct {
	Height    int64           `json:"height"`
	TxHash    string          `json:"txhash"`
	Code      uint32          `json:"code,omitempty"`
	Data      string          `json:"data,omitempty"`
	RawLog    string          `json:"raw_log,omitempty"`
	Logs      ABCIMessageLogs `json:"logs,omitempty"`
	Info      string          `json:"info,omitempty"`
	GasWanted int64           `json:"-"`
	GasUsed   int64           `json:"-"`
	Tags      StringTags      `json:"tags,omitempty"`
	Codespace string          `json:"codespace,omitempty"`
	Tx        Tx              `json:"tx,omitempty"`
	Timestamp string          `json:"timestamp,omitempty"`
}

// ABCIMessageLogs represents a slice of ABCIMessageLog.
type ABCIMessageLogs []ABCIMessageLog

// ABCIMessageLog defines a structure containing an indexed tx ABCI message log.
type ABCIMessageLog struct {
	MsgIndex int    `json:"msg_index"`
	Success  bool   `json:"success"`
	Log      string `json:"log"`
}

// A slice of StringTag
type StringTags []StringTag

// A KVPair where the Key and Value are both strings, rather than []byte
type StringTag struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}