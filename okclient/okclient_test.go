package okclient

import (
	"errors"
	"fmt"
	"github.com/ok-chain/gosdk/common/queryParams"
	"github.com/ok-chain/gosdk/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"testing"
)

const (
	RPC_URL = "localhost:26657"

)

func TestNewClient(t *testing.T) {
	okCli := NewClient(RPC_URL)

	accountParams := queryParams.NewQueryAccTokenParams("","all")


	jsonBytes, err := okCli.cdc.MarshalJSON(accountParams)
	assertNotEqual(t, err, nil)

	//fmt.Println(jsonBytes)
	path := "custom/token/accounts/okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	opts := rpcclient.ABCIQueryOptions{
		Height: 0,
		Prove:  false,
	}
	result, err := okCli.cli.ABCIQueryWithOptions(path, jsonBytes, opts)
	assertNotEqual(t, err, nil)
	//fmt.Println(result)
	resp := result.Response
	if !resp.IsOK() {
		t.Error(errors.New(resp.Log))
	}

	var accountResponse types.AccountTokensInfo
	if err = okCli.cdc.UnmarshalJSON(resp.Value, &accountResponse); err != nil {
		assertNotEqual(t, err, nil)
	}

	fmt.Println(accountResponse)

}
