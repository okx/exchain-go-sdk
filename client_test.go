package gosdk

//
//// just a temporary test, it will be removed later

//// need test
//
//func TestQueryCandles(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//	candles, err := client.Backend().QueryCandles("tbtc-44f_tusdk-0cd", 60, 100)
//	require.NoError(t, err)
//	for _, line := range candles {
//		fmt.Println(line)
//	}
//}
//
//func TestQueryTickers(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//	tickers, err := client.Backend().QueryTickers(10)
//	require.NoError(t, err)
//	for _, t := range tickers {
//		fmt.Println(t)
//	}
//}
//
//func TestQueryRecentTxRecord(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//	record, err := client.Backend().QueryRecentTxRecord("tbtc-44f_tusdk-0cd", 0, int(time.Now().Unix()), 0, 10)
//	require.NoError(t, err)
//	for _, res := range record {
//		fmt.Println(res)
//	}
//}
//
//func TestQueryOpenOrders(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//
//	product := "tbtc-44f_tusdk-0cd"
//	side := "BUY"
//	start, end := 1, int(time.Now().Unix())
//	page, perPage := 0, 10
//
//	openOrdersList, err := client.Backend().QueryOpenOrders(backendAddr, product, side, start, end, page, perPage)
//	require.NoError(t, err)
//	for _, order := range openOrdersList {
//		fmt.Println(order)
//	}
//}
//
//func TestQueryClosedOrders(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//
//	product := "tbtc-44f_tusdk-0cd"
//	side := "BUY"
//	start, end := 1, int(time.Now().Unix())
//	page, perPage := 0, 10
//
//	closedOrdersList, err := client.Backend().QueryClosedOrders(backendAddr, product, side, start, end, page, perPage)
//	require.NoError(t, err)
//	for _, order := range closedOrdersList {
//		fmt.Println(order)
//	}
//}
//
//func TestQueryDeals(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//
//	product := "tbtc-44f_tusdk-0cd"
//	side := "BUY"
//	start, end := 1, int(time.Now().Unix())
//	page, perPage := 0, 10
//
//	deals, err := client.Backend().QueryDeals(backendAddr, product, side, start, end, page, perPage)
//	require.NoError(t, err)
//	for _, deal := range deals {
//		fmt.Println(deal)
//	}
//}
//
//func TestQueryTransactions(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
//	client := NewClient(config)
//
//	typeCode := 0
//	start, end := 1, int(time.Now().Unix())
//	page, perPage := 0, 10
//
//	txs, err := client.Backend().QueryTransactions(backendAddr, typeCode, start, end, page, perPage)
//	require.NoError(t, err)
//	for _, tx := range txs {
//		fmt.Println(tx)
//	}
//}

//func TestQueryTxsResult(t *testing.T) {
//	config := NewClientConfig("tcp://192.168.13.130:20057", BroadcastBlock)
//	client := NewClient(config)
//	// get searching string
//	searchStr := `message.sender=okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya`
//	txsRes, err := client.Tendermint().QueryTxsResult(searchStr, 1, 30)
//	require.NoError(t, err)
//	fmt.Printf("%+v\n", txsRes)
//}
//

//config, _ := NewClientConfig("tcp://127.0.0.1:10057", "", BroadcastBlock, "0.01okt", 0)
//client := NewClient(config)
