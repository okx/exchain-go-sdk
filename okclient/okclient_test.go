package okclient

import (
	"fmt"
	"testing"
)

const (
	RPC_URL = "localhost:26657"
)

func TestNewClient(t *testing.T) {
	okCli := NewClient(RPC_URL)
	fmt.Println(okCli)

	//accountParam := queryParams.AccountParam{
	//	Symbol: "",
	//	Show:   "all",
	//}
	//
	//jsonBytes, err := okCli.cdc.MarshalJSON(accountParam)
	//assertEqual(t, err, nil)
	//
	//path := "custom/token/accounts/okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	//opts := rpcCli.ABCIQueryOptions{
	//	Height: 0,
	//	Prove:  false,
	//}
	//result, err := okCli.cli.ABCIQueryWithOptions(path, jsonBytes, opts)
	//assertEqual(t, err, nil)
	//resp := result.Response
	//if !resp.IsOK() {
	//	t.Error(errors.New(resp.Log))
	//}
	//
	//var accountResponse token.AccountResponse
	//if err = okCli.cdc.UnmarshalJSON(resp.Value, &accountResponse); err != nil {
	//	assertEqual(t, err, nil)
	//}
	//
	//fmt.Println(accountResponse)

}

func assertEqual(t *testing.T, err, b interface{}) {
	if err != b {
		t.Errorf("test failed: %s", err)
	}
}
