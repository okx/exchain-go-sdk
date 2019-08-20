package okclient

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueryABCIInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryABCIInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryConsenusState(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryConsenusState()
	assertNotEqual(t, err, nil)
	fmt.Println(string(resp))
}

func TestQueryDumpConsenusState(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryDumpConsenusState()
	assertNotEqual(t, err, nil)
	fmt.Println(string(resp))
}

func TestQueryNetInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryNetInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryGenesisFile(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryGenesisFile()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryHealthInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryHealthInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryUnconfirmedTxsNum(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryUnconfirmedTxsNum(30)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryStateInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryStateInfo()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryABCITokenpair(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryABCITokenpair()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryBlock(t *testing.T) {
	okCli := NewClient(rpcUrl)
	var height int64 = 1024
	resp, err := okCli.QueryBlock(&height)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryBlockResults(t *testing.T) {
	okCli := NewClient(rpcUrl)
	// input nil means query the latest block info
	resp, err := okCli.QueryBlockResults(nil)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryBlockchainInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	resp, err := okCli.QueryBlockchainInfo(0, 10)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryCommit(t *testing.T) {
	okCli := NewClient(rpcUrl)
	var height int64 = 1024
	resp, err := okCli.QueryCommit(&height)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryTx(t *testing.T) {
	okCli := NewClient(rpcUrl)
	// get tx hash bytes
	txHash, err := hex.DecodeString("12CF714D13D9B86EDCCBE41BF55845BF96613977AFF8E503C5A5349A50841F9A")
	assertNotEqual(t, err, nil)
	resp, err := okCli.QueryTx(txHash, true)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryTxOnHeight(t *testing.T) {
	okCli := NewClient(rpcUrl)
	var queryStr  ="tx.height=202996"
	resp, err := okCli.QueryTxOnHeight(queryStr, true,1,30)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}
