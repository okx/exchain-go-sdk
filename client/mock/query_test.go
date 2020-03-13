package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/okex/okchain-go-sdk/client"
	"github.com/okex/okchain-go-sdk/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"testing"
	"time"
)

const (
	addr = "okchain1h643csgkdr8yzey22tekxwzjlxgn87sm0ust5u"
)

func TestGetAccountInfoByAddr(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	accAddress, err := types.AccAddressFromBech32(addr)
	accPubKey, err := client.GetAccPubKeyBech32("okchainpub1addwnpepqwqz3e53lu43pf9zhpxyxrwl9zpzrt3l4y2fsj9dqq947p604z6wzs7u070")
	accounts := types.BaseAccount{
		Address: accAddress,
		Coins: types.Coins{
			types.Coin{
				Denom:  "okt",
				Amount: types.NewInt(100000),
			},
			types.Coin{
				Denom:  "acoin",
				Amount: types.NewInt(2000),
			},
		},
		PubKey:        accPubKey,
		AccountNumber: 2917,
		Sequence:      15,
	}

	accountBytes, err := cli.GetCdc().MarshalBinaryBare(accounts)
	assert.Equal(t, err, nil)

	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = accountBytes

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	acc, err := cli.GetAccountInfoByAddr(addr)
	assert.Equal(t, err, nil)
	fmt.Println(acc)
}

func TestGetTokensInfoByAddr(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokensInfo := `{"address":"okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya","currencies":[{"symbol":"acoin","available":"10000000.00000000","freeze":"0","locked":"0"},{"symbol":"bcoin","available":"10000000.00000000","freeze":"0","locked":"0"}]}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(tokensInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	tokensInfoResult, err := cli.GetTokensInfoByAddr(addr)
	assert.Equal(t, err, nil)
	fmt.Println(tokensInfoResult)
}

func TestGetTokenInfoByAddr(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokenInfo := `{"address":"okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya","currencies":[{"symbol":"okt","available":"10000000.00000000","freeze":"0","locked":"0"}]}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(tokenInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	tokenInfoResult, err := cli.GetTokenInfoByAddr(addr, "okt")
	assert.Equal(t, err, nil)
	fmt.Println(tokenInfoResult)
}

func TestGetTokensInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokensInfo := `[{"desc":"test testcoin001","symbol":"coin01-7f3","originalSymbol":"coin01","wholeName":"testcoin01","totalSupply":"1000000","owner":"okchain1sv8shwdyyf5hxadawe4nw9uc82lxzk4zymvc4z","mintable":true},{"desc":"k8s blockchain coin","symbol":"k8coin-a57","originalSymbol":"k8coin","wholeName":"k8coin","totalSupply":"1199990","owner":"okchain1h643csgkdr8yzey22tekxwzjlxgn87sm0ust5u","mintable":true},{"desc":"Bitcoin in testnet，1:1 anchoring with Bitcoin","symbol":"tbtc","originalSymbol":"tbtc","wholeName":"Testnet Bitcoin","totalSupply":"21000000","owner":"okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya","mintable":false}]`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(tokensInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	tokensInfoResult, err := cli.GetTokensInfo()
	assert.Equal(t, err, nil)
	for _, t := range tokensInfoResult {
		fmt.Println(t)
	}
}

func TestGetTokenInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tokensInfo := `{"desc":"OKChain token in testnet","symbol":"tokt","originalSymbol":"TOKT","wholeName":"Testnet OK token","totalSupply":"1000000000","owner":"okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya","mintable":true}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(tokensInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	tokenInfo, err := cli.GetTokenInfo("tokt")
	assert.Equal(t, err, nil)
	fmt.Println(tokenInfo)
}

func TestGetProductsInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productsInfo := `[{"baseAssetSymbol":"tokt","quoteAssetSymbol":"tusdk","price":"2.82600000","maxPriceDigit":"4","maxSizeDigit":"4","minTradeSize":"0.00010000","tokenPairId":"1"},{"baseAssetSymbol":"tokb","quoteAssetSymbol":"tusdk","price":"2.91700000","maxPriceDigit":"4","maxSizeDigit":"4","minTradeSize":"0.00010000","tokenPairId":"2"},{"baseAssetSymbol":"tbtc","quoteAssetSymbol":"tusdk","price":"7887.60000000","maxPriceDigit":"4","maxSizeDigit":"4","minTradeSize":"0.10000000","tokenPairId":"3"},{"baseAssetSymbol":"zcoin1-d7c","quoteAssetSymbol":"tusdk","price":"1.00000000","maxPriceDigit":"7","maxSizeDigit":"1","minTradeSize":"0.00010000","tokenPairId":"4"},{"baseAssetSymbol":"zcoin1-d7c","quoteAssetSymbol":"tokt","price":"1.00000000","maxPriceDigit":"7","maxSizeDigit":"1","minTradeSize":"0.00010000","tokenPairId":"5"},{"baseAssetSymbol":"tusdk","quoteAssetSymbol":"zcoin1-d7c","price":"1.00000000","maxPriceDigit":"7","maxSizeDigit":"1","minTradeSize":"0.00010000","tokenPairId":"6"}]`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(productsInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	productsList, err := cli.GetProductsInfo()
	assert.Equal(t, err, nil)
	for _, p := range productsList {
		fmt.Println(p)
	}
}

func TestGetDepthbookInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	depthBookInfo := `{"asks":[{"price":"10.01","quantity":"1000"}],"bids":[{"price":"10.5","quantity":"200"}]}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(depthBookInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	depthBook, err := cli.GetDepthbookInfo("k8coin_tokt")
	assert.Equal(t, err, nil)
	for _, ask := range depthBook.Asks {
		fmt.Println(ask)
	}
	for _, bid := range depthBook.Bids {
		fmt.Println(bid)
	}
}

func TestGetCandlesInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	depthBookInfo := `{"code":0,"data":[["2018-07-12T04:00:00.000Z","6343.3587","6345.0453","6142.2336","6186.8354","8429.75582698"]],"detail_msg":"string","msg":"string"}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(depthBookInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	candles, err := cli.GetCandlesInfo("k8coin_tokt", 60, 100)
	assert.Equal(t, err, nil)
	for _, line := range candles {
		fmt.Println(line)
	}
}

func TestGetTickersInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tickersInfo := `{"code":0,"data":[{"close":"29.777","high":"55.44","low":"22.22","open":"55.44","price":"29.777","product":"bcoin-2ac_okt","symbol":"bcoin-2ac_okt","timestamp":"2019-07-25T09:49:04.954Z","volume":"266.64"}],"detail_msg":"","msg":""}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(tickersInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	tickers, err := cli.GetTickersInfo(10)
	assert.Equal(t, err, nil)
	for _, ticker := range tickers {
		fmt.Println(ticker)
	}
}

func TestGetRecentTxRecord(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recentTxRecordInfo := `{"code":0,"msg":"","detailMsg":"","data":{"data":[{"timestamp":1559790137,"blockHeight":386355,"product":"acoin-564_okt","price":3,"volume":0.25},{"timestamp":1559789554,"blockHeight":386159,"product":"acoin-564_okt","price":1.9999,"volume":2.9999},{"timestamp":1559788804,"blockHeight":385931,"product":"acoin-564_okt","price":1,"volume":1}],"paramPage":{"page":1,"perPage":50,"total":3}}}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(recentTxRecordInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	records, err := cli.GetRecentTxRecord("k8coin_tokt", 0, int(time.Now().Unix()), 0, 10)
	assert.Equal(t, err, nil)
	for _, record := range records {
		fmt.Println(record)
	}
}

func TestGetOpenOrders(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	openOrdersInfo := `{"code":0,"msg":"","detailMsg":"","data":{"data":[{"txhash":"2144D0F85B67D9508066004400BF8044010ED5FC4B43417F9A44CDC3EBAD9765","order_id":"O0000000008-0000","sender":"cosmos1hghms6dtm8quxegrkcnw4wnzj5e5sc4am0gxyr","product":"k8coin_tokt","side":"BUY","price":"10.000000000000000000","quantity":"1.100000000000000000","status":0,"filled_avg_price":"10.000000000000000000","remain_quantity":"0.100000000000000000","timestamp":1553842734}],"paramPage":{"total":"1","page":1,"perPage":10}}}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(openOrdersInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	product := "k8coin_tokt"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	openOrdersList, err := cli.GetOpenOrders(addr, product, side, start, end, page, perPage)
	assert.Equal(t, err, nil)
	for _, order := range openOrdersList {
		fmt.Println(order)
	}
}

func TestGetClosedOrders(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	closedOrdersInfo := `{"code":0,"msg":"","detailMsg":"","data":{"data":[{"txhash":"2144D0F85B67D9508066004400BF8044010ED5FC4B43417F9A44CDC3EBAD9765","order_id":"O0000000008-0000","sender":"cosmos1hghms6dtm8quxegrkcnw4wnzj5e5sc4am0gxyr","product":"k8coin_tokt","side":"BUY","price":"10.000000000000000000","quantity":"1.100000000000000000","status":0,"filled_avg_price":"10.000000000000000000","remain_quantity":"0.100000000000000000","timestamp":1553842734}],"paramPage":{"total":"1","page":1,"perPage":10}}}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(closedOrdersInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	product := "k8coin_tokt"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	openOrdersList, err := cli.GetClosedOrders(addr, product, side, start, end, page, perPage)
	assert.Equal(t, err, nil)
	for _, order := range openOrdersList {
		fmt.Println(order)
	}
}

func TestGetDealsInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dealsInfo := `{"code":0,"msg":"","detailMsg":"","data":{"data":[{"timestamp":1558407585,"blockHeight":463,"orderId":"ID0000000463-1","sender":"okchain15wv9q08rv0f8dg08scv2ps45hs6v8qx37466qj","product":"k8coin_tokt","side":"BUY","price":10,"volume":1,"fee":"0.00400000okt"},{"timestamp":1558407585,"blockHeight":463,"orderId":"ID0000000010-1","sender":"okchain1lzekrp7dezrs940m7c0nnhjvyhlzppnaf6vjsy","product":"k8coin_tokt","side":"BUY","price":10,"volume":1,"fee":"0.00400000okt"}],"paramPage":{"page":1,"perPage":10,"total":2}}}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(dealsInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	product := "k8coin_tokt"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	dealsInfoList, err := cli.GetDealsInfo(addr, product, side, start, end, page, perPage)
	assert.Equal(t, err, nil)
	for _, deal := range dealsInfoList {
		fmt.Println(deal)
	}
}

func TestGetTransactionsInfo(t *testing.T) {
	cli := client.NewClient("rpcUrl")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	transactionsInfo := `{"code":0,"msg":"","detailMsg":"","data":{"data":[{"txHash":"3BEE2A0FDDD5EB077236879E139DC565580139F61ED6E391B2557D4A8F74BE83","type":1,"address":"okchain1lzekrp7dezrs940m7c0nnhjvyhlzppnaf6vjsy","symbol":"tokt","side":3,"quantity":"1.00000000","fee":"0.01250000okt","timestamp":1558407348}],"paramPage":{"page":1,"perPage":10,"total":1}}}`
	var result ctypes.ResultABCIQuery
	result.Response.Code = 0
	result.Response.Value = []byte(transactionsInfo)

	mockClient := NewMockClient(ctrl)
	mockClient.EXPECT().ABCIQueryWithOptions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&result, nil)
	cli.SetClient(mockClient)

	type_ := 1
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	transactionsInfoList, err := cli.GetTransactionsInfo(addr, type_, start, end, page, perPage)
	assert.Equal(t, err, nil)
	for _, tx := range transactionsInfoList {
		fmt.Println(tx)
	}
}
