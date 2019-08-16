package okclient

import (
	"fmt"
	"testing"
)

var (
	name     = "alice"
	passWd   = "12345678"
	mnemonic = "sustain hole urban away boy core lazy brick wait drive tiger tell"
	addr     = "okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	rpcUrl   = "localhost:26657"
)

func TestGetAccountInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	acc, err := okCli.GetAccountInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(acc)

}

func assertNotEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("test failed: %s", a)
	}
}

func assertEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("test failed: %s", a)
	}
}
