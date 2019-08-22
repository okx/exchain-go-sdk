package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueryABCIInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryABCIInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryConsenusState(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryConsenusState()
	assertNotEqual(t, err, nil)
	fmt.Println(string(resp))
}

func TestQueryDumpConsenusState(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryDumpConsenusState()
	assertNotEqual(t, err, nil)
	fmt.Println(string(resp))
}

func TestQueryNetInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryNetInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryGenesisFile(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryGenesisFile()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryHealthInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryHealthInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryUnconfirmedTxsNum(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryUnconfirmedTxsNum(30)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryStateInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryStateInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryABCITokenpair(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryABCITokenpair()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryBlock(t *testing.T) {
	cli := NewClient(rpcUrl)
	var height int64 = 1024
	resp, err := cli.QueryBlock(&height)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryBlockResults(t *testing.T) {
	cli := NewClient(rpcUrl)
	// input nil means query the latest block info
	resp, err := cli.QueryBlockResults(nil)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryBlockchainInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryBlockchainInfo(0, 10)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryCommit(t *testing.T) {
	cli := NewClient(rpcUrl)
	var height int64 = 1024
	resp, err := cli.QueryCommit(&height)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryTx(t *testing.T) {
	cli := NewClient(rpcUrl)
	// get tx hash bytes
	txHash, err := hex.DecodeString("12CF714D13D9B86EDCCBE41BF55845BF96613977AFF8E503C5A5349A50841F9A")
	assertNotEqual(t, err, nil)
	resp, err := cli.QueryTx(txHash, true)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryTxOnHeight(t *testing.T) {
	cli := NewClient(rpcUrl)
	var queryStr = "tx.height=202996"
	resp, err := cli.QueryTxOnHeight(queryStr, true, 1, 30)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryCurrentValidators(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryCurrentValidators()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}
