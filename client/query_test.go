package client

import (
	"fmt"
	"testing"
	"time"
)

const (
	rpcUrl = "tcp://127.0.0.1:10057"
)










func TestGetTransactionsInfo(t *testing.T) {
	cli := NewClient(rpcUrl)

	type_ := 1
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	transactionsInfo, err := cli.GetTransactionsInfo(addr, type_, start, end, page, perPage)
	assertNotEqual(t, err, nil)
	for _, tx := range transactionsInfo {
		fmt.Println(tx)
	}
}


func assertNotEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("test failed: %s", a)
	}
}

