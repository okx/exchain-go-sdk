package client

import (
	"fmt"
	"testing"
	"time"
)

const (
	addr   = "okchain1mm43akh88a3qendlmlzjldf8lkeynq68r8l6ts"
	rpcUrl = "3.13.150.20:26657"

)

func TestGetAccountInfoByAddr(t *testing.T) {
	cli := NewClient(rpcUrl)
	acc, err := cli.GetAccountInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(acc)
}

func TestGetTokensInfoByAddr(t *testing.T) {
	cli := NewClient(rpcUrl)
	tokensInfo, err := cli.GetTokensInfoByAddr(addr)
	assertNotEqual(t, err, nil)
	fmt.Println(tokensInfo)
}

func TestGetTokenInfoByAddr(t *testing.T) {
	cli := NewClient(rpcUrl)
	tokenInfo, err := cli.GetTokenInfoByAddr(addr, "tokt")
	assertNotEqual(t, err, nil)
	fmt.Println(tokenInfo)
}

func TestGetTokensInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	tokensInfo, err := cli.GetTokensInfo()
	assertNotEqual(t, err, nil)
	for _, t := range tokensInfo {
		fmt.Println(t)
	}
}

func TestGetTokenInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	tokenInfo, err := cli.GetTokenInfo("tokt")
	assertNotEqual(t, err, nil)
	fmt.Println(tokenInfo)
}

func TestGetProductsInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	productsList, err := cli.GetProductsInfo()
	assertNotEqual(t, err, nil)
	for _, p := range productsList {
		fmt.Println(p)
	}
}

func TestGetDepthbookInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	depthbook, err := cli.GetDepthbookInfo("xxb_tokt")
	assertNotEqual(t, err, nil)
	for _, ask := range depthbook.Asks {
		fmt.Println(ask)
	}
	for _, bid := range depthbook.Bids {
		fmt.Println(bid)
	}

}

func TestGetCandlesInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	candles, err := cli.GetCandlesInfo("xxb_tokt", 60, 100)
	assertNotEqual(t, err, nil)
	for _, line := range candles {
		fmt.Println(line)
	}
}

func TestGetTickersInfo(t *testing.T) {
	cli := NewClient(rpcUrl)
	tickers, err := cli.GetTickersInfo(10)
	assertNotEqual(t, err, nil)
	for _, ticker := range tickers {
		fmt.Println(ticker)
	}
}

func TestGetRecentTxRecord(t *testing.T) {
	cli := NewClient(rpcUrl)
	records, err := cli.GetRecentTxRecord("xxb_okb", 0, int(time.Now().Unix()), 0, 10)
	assertNotEqual(t, err, nil)
	for _, record := range records {
		fmt.Println(record)
	}
}

func TestGetOpenOrders(t *testing.T) {
	cli := NewClient(rpcUrl)

	product := "xxb_okb"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	openOrdersList, err := cli.GetOpenOrders(addr, product, side, start, end, page, perPage)
	assertNotEqual(t, err, nil)
	for _, order := range openOrdersList {
		fmt.Println(order)
	}
}

func TestGetClosedOrders(t *testing.T) {
	cli := NewClient(rpcUrl)

	product := "xxb_okb"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	openOrdersList, err := cli.GetClosedOrders(addr, product, side, start, end, page, perPage)
	assertNotEqual(t, err, nil)
	for _, order := range openOrdersList {
		fmt.Println(order)
	}
}

func TestGetDealsInfo(t *testing.T) {
	cli := NewClient(rpcUrl)

	product := "xxb_okb"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	dealsInfo, err := cli.GetDealsInfo(addr, product, side, start, end, page, perPage)
	assertNotEqual(t, err, nil)
	for _, deal := range dealsInfo {
		fmt.Println(deal)
	}
}

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

func assertEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("test failed: %s", a)
	}
}
