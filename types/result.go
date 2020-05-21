package types

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Result is the union of ResponseFormat and ResponseCheckTx
type Result struct {
	// Code is the response code, is stored back on the chain
	Code CodeType

	// Codespace is the string referring to the domain of an error
	Codespace CodespaceType

	// Data is any data returned from the app
	// Data has to be length prefixed in order to separate
	// results from multiple msgs executions
	Data []byte

	// Log contains the txs log information. NOTE: nondeterministic
	Log string

	// GasWanted is the maximum units of work we allow this tx to perform
	GasWanted uint64

	// GasUsed is the amount of gas actually consumed. NOTE: unimplemented
	GasUsed uint64

	// Events contains a slice of Event objects that were emitted during some execution
	Events Events
}

// IsOK shows whether the result is successful
// TODO: In the future, more codes may be OK
func (res Result) IsOK() bool {
	return res.Code.IsOK()
}

// ABCIMessageLogs represents a slice of ABCIMessageLog
type ABCIMessageLogs []ABCIMessageLog

// String implements the fmt.Stringer interface for the ABCIMessageLogs type
func (logs ABCIMessageLogs) String() (str string) {
	if logs != nil {
		raw, err := Cdc.MarshalJSON(logs)
		if err == nil {
			str = string(raw)
		}
	}

	return str
}

// ABCIMessageLog defines a structure containing an indexed tx ABCI message log
type ABCIMessageLog struct {
	MsgIndex uint16 `json:"msg_index"`
	Success  bool   `json:"success"`
	Log      string `json:"log"`

	// Events contains a slice of Event objects that were emitted during some execution
	Events StringEvents `json:"events"`
}

// TxResponse defines a structure containing relevant tx data and metadata
// The tags are stringified and the log is JSON decoded
type TxResponse struct {
	Height    int64           `json:"height,omitempty"`
	TxHash    string          `json:"txhash"`
	Code      uint32          `json:"code,omitempty"`
	Data      string          `json:"data,omitempty"`
	RawLog    string          `json:"raw_log,omitempty"`
	Logs      ABCIMessageLogs `json:"logs,omitempty"`
	Info      string          `json:"info,omitempty"`
	GasWanted int64           `json:"-"`
	GasUsed   int64           `json:"-"`
	Codespace string          `json:"codespace,omitempty"`
	Tx        Tx              `json:"tx,omitempty"`
	Timestamp string          `json:"timestamp,omitempty"`

	// DEPRECATED: Remove in the next next major release in favor of using the ABCIMessageLog.Events field
	Events StringEvents `json:"events,omitempty"`
}

// NewResponseFormatBroadcastTxCommit returns a TxResponse given a ResultBroadcastTxCommit from tendermint
func NewResponseFormatBroadcastTxCommit(res *ctypes.ResultBroadcastTxCommit) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	if !res.CheckTx.IsOK() {
		return newTxResponseCheckTx(res)
	}

	return newTxResponseDeliverTx(res)
}

func newTxResponseCheckTx(res *ctypes.ResultBroadcastTxCommit) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	var txHash string
	if res.Hash != nil {
		txHash = res.Hash.String()
	}

	parsedLogs, err := ParseABCILogs(res.CheckTx.Log)
	if err != nil {
		log.Println(err)
	}

	return TxResponse{
		Height:    res.Height,
		TxHash:    txHash,
		Code:      res.CheckTx.Code,
		Data:      strings.ToUpper(hex.EncodeToString(res.CheckTx.Data)),
		RawLog:    res.CheckTx.Log,
		Logs:      parsedLogs,
		Info:      res.CheckTx.Info,
		GasWanted: res.CheckTx.GasWanted,
		GasUsed:   res.CheckTx.GasUsed,
		Events:    StringifyEvents(res.CheckTx.Events),
		Codespace: res.CheckTx.Codespace,
	}
}

func newTxResponseDeliverTx(res *ctypes.ResultBroadcastTxCommit) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	var txHash string
	if res.Hash != nil {
		txHash = res.Hash.String()
	}

	parsedLogs, _ := ParseABCILogs(res.DeliverTx.Log)

	return TxResponse{
		Height:    res.Height,
		TxHash:    txHash,
		Code:      res.DeliverTx.Code,
		Data:      strings.ToUpper(hex.EncodeToString(res.DeliverTx.Data)),
		RawLog:    res.DeliverTx.Log,
		Logs:      parsedLogs,
		Info:      res.DeliverTx.Info,
		GasWanted: res.DeliverTx.GasWanted,
		GasUsed:   res.DeliverTx.GasUsed,
		Events:    StringifyEvents(res.DeliverTx.Events),
		Codespace: res.DeliverTx.Codespace,
	}
}

// NewResponseFormatBroadcastTx returns a TxResponse given a ResultBroadcastTx from tendermint
func NewResponseFormatBroadcastTx(res *ctypes.ResultBroadcastTx) TxResponse {
	if res == nil {
		return TxResponse{}
	}

	parsedLogs, err := ParseABCILogs(res.Log)
	if err != nil {
		log.Println(err)
	}

	return TxResponse{
		Code:   res.Code,
		Data:   res.Data.String(),
		RawLog: res.Log,
		Logs:   parsedLogs,
		TxHash: res.Hash.String(),
	}
}

// String returns a human readable string representation of TxResponse
func (r TxResponse) String() string {
	var sb strings.Builder
	if _, err := sb.WriteString("Response:\n"); err != nil {
		log.Println(err)
	}

	if r.Height > 0 {
		if _, err := sb.WriteString(fmt.Sprintf("  Height: %d\n", r.Height)); err != nil {
			log.Println(err)
		}
	}

	if r.TxHash != "" {
		if _, err := sb.WriteString(fmt.Sprintf("  TxHash: %s\n", r.TxHash)); err != nil {
			log.Println(err)
		}
	}

	if r.Code > 0 {
		if _, err := sb.WriteString(fmt.Sprintf("  Code: %d\n", r.Code)); err != nil {
			log.Println(err)
		}
	}

	if r.Data != "" {
		if _, err := sb.WriteString(fmt.Sprintf("  Data: %s\n", r.Data)); err != nil {
			log.Println(err)
		}
	}

	if r.RawLog != "" {
		if _, err := sb.WriteString(fmt.Sprintf("  Raw Log: %s\n", r.RawLog)); err != nil {
			log.Println(err)
		}
	}

	if r.Logs != nil {
		if _, err := sb.WriteString(fmt.Sprintf("  Logs: %s\n", r.Logs)); err != nil {
			log.Println(err)
		}

	}

	if r.Info != "" {
		if _, err := sb.WriteString(fmt.Sprintf("  Info: %s\n", r.Info)); err != nil {
			log.Println(err)
		}
	}

	//if r.GasWanted != 0 {
	//	sb.WriteString(fmt.Sprintf("  GasWanted: %d\n", r.GasWanted))
	//}
	//
	//if r.GasUsed != 0 {
	//	sb.WriteString(fmt.Sprintf("  GasUsed: %d\n", r.GasUsed))
	//}

	if r.Codespace != "" {
		if _, err := sb.WriteString(fmt.Sprintf("  Codespace: %s\n", r.Codespace)); err != nil {
			log.Println(err)
		}
	}

	if r.Timestamp != "" {
		if _, err := sb.WriteString(fmt.Sprintf("  Timestamp: %s\n", r.Timestamp)); err != nil {
			log.Println(err)
		}
	}

	if len(r.Events) > 0 {
		if _, err := sb.WriteString(fmt.Sprintf("  Events: \n%s\n", r.Events.String())); err != nil {
			log.Println(err)
		}
	}

	return strings.TrimSpace(sb.String())
}

// Empty returns true if the response is empty
func (r TxResponse) Empty() bool {
	return r.TxHash == "" && r.Logs == nil
}

// SearchTxsResult defines a structure for querying txs pageable
type SearchTxsResult struct {
	TotalCount int          `json:"total_count"` // Count of all txs
	Count      int          `json:"count"`       // Count of txs in current page
	PageNumber int          `json:"page_number"` // Index of current page, start from 1
	PageTotal  int          `json:"page_total"`  // Count of total pages
	Limit      int          `json:"limit"`       // Max count txs per page
	Txs        []TxResponse `json:"txs"`         // List of txs in current page
}

// ParseABCILogs attempts to parse a stringified ABCI tx log into a slice of ABCIMessageLog types
// It returns an error upon JSON decoding failure
func ParseABCILogs(logs string) (res ABCIMessageLogs, err error) {
	err = json.Unmarshal([]byte(logs), &res)
	return res, err
}
