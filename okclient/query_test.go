package okclient

import (
	"fmt"
	"testing"
)

const (
	addr   = "okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	rpcUrl = "localhost:26657"
)

func TestGetAccountInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	acc, err := okCli.GetAccountInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(acc)

}

func TestGetTokensInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokensInfo, err := okCli.GetTokensInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(tokensInfo)
}

func TestGetTokenInfoByAddr(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokenInfo, err := okCli.GetTokenInfoByAddr(addr, "okb")
	assertNotEqual(t, err, nil)
	fmt.Println(tokenInfo)
}

func TestGetTokensInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokensInfo, err := okCli.GetTokensInfo()
	assertNotEqual(t, err, nil)
	for _, t := range tokensInfo {
		fmt.Println(t)
	}
}

func TestGetTokenInfo(t *testing.T) {
	okCli := NewClient(rpcUrl)
	tokenInfo, err := okCli.GetTokenInfo("okb")
	assertNotEqual(t, err, nil)
	fmt.Println(tokenInfo)
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
