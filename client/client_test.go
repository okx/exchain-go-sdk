package client

import (
	"errors"
	"fmt"
	"github.com/okex/okchain-go-sdk/common/queryParams"
	"github.com/okex/okchain-go-sdk/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"testing"
)


func TestNewClient(t *testing.T) {
	cli := NewClient(rpcUrl)

	accountParams := queryParams.NewQueryAccTokenParams("","all")


	jsonBytes, err := cli.cdc.MarshalJSON(accountParams)
	assertNotEqual(t, err, nil)

	//fmt.Println(jsonBytes)
	path := "custom/token/accounts/okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	opts := rpcclient.ABCIQueryOptions{
		Height: 0,
		Prove:  false,
	}
	result, err := cli.cli.ABCIQueryWithOptions(path, jsonBytes, opts)
	assertNotEqual(t, err, nil)
	resp := result.Response
	if !resp.IsOK() {
		t.Error(errors.New(resp.Log))
	}

	var accountResponse types.AccountTokensInfo
	if err = cli.cdc.UnmarshalJSON(resp.Value, &accountResponse); err != nil {
		assertNotEqual(t, err, nil)
	}

	fmt.Println(accountResponse)

}
