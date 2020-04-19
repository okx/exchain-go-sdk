package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestQueryProposals(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryProposals()
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}

func TestQueryProposalByID(t *testing.T) {
	cli := NewClient(rpcUrl)
	resp, err := cli.QueryProposalByID(1)
	assertNotEqual(t, err, nil)
	jsonBytes, err := json.Marshal(resp)
	assertNotEqual(t, err, nil)
	fmt.Println(string(jsonBytes))
}
