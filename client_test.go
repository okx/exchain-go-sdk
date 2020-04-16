package sdk

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// just a temporary test, it will be removed later

const (
	name   = "alice"
	passWd = "12345678"
	// sender's mnemonic
	mnemonic = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	addr     = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	// target mnemonic
	targetMnemonic = "pepper basket run install fury scheme journey worry tumble toddler swap change"
	targetAddr     = "okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7"

	backendAddr = "okchain10q0rk5qnyag7wfvvt7rtphlw589m7frsmyq4ya"
)

// transact tx

func TestDelegate(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	resp, err := client.Staking().Delegate(fromInfo, passWd, "1024.1024okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestUnbond(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	resp, err := client.Staking().Unbond(fromInfo, passWd, "0.1024okt", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestVote(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	valsToVoted := []string{"okchainvaloper1n62v94azspas83uucwxg347jqmfma90fwx7nxt", "okchainvaloper1wsrrv0q4ldqjm2lxayuscwthcht55crdnt6her"}
	resp, err := client.Staking().Vote(fromInfo, passWd, valsToVoted, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestDestroyValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(
		"novel tomorrow scorpion cross immense photo wrap acquire midnight about what clean", name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	resp, err := client.Staking().DestroyValidator(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestCreateValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	pubkeyStr := "okchainvalconspub1zcjduepqghrtvkngejwese62wg49ewskz4r93vkyj3md5mg5rf7twcc6jduqpqw66q"
	resp, err := client.Staking().CreateValidator(fromInfo, passWd, pubkeyStr, "my moniker", "my identity",
		"my website", "my details", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)
}

func TestEditValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(
		"ready edge sketch vibrant cause snake donor trophy cruise pulse vanish siren", name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	resp, err := client.Staking().EditValidator(fromInfo, passWd, "my moniker", "my identity", "my website",
		"my details", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(resp)

}

func TestRegisterProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	sequence := accInfo.GetSequence()
	res, err := client.Staking().Delegate(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	sequence++
	res, err = client.Staking().RegisterProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)
	fmt.Println(res)
}

func TestUnregisterProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(addr)
	require.NoError(t, err)

	res, err := client.Staking().UnregisterProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestBindProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	valMnemo := "ready edge sketch vibrant cause snake donor trophy cruise pulse vanish siren"
	// validator becomes a proxy
	valAcc, _, err := utils.CreateAccountWithMnemo(valMnemo, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(valAcc.GetAddress().String())
	require.NoError(t, err)

	sequence := accInfo.GetSequence()
	res, err := client.Staking().Delegate(valAcc, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	sequence++
	res, err = client.Staking().RegisterProxy(valAcc, passWd, "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	// delegator tries to bind proxy
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err = client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	sequence = accInfo.GetSequence()
	res, err = client.Staking().Delegate(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)

	sequence++
	res, err = client.Staking().BindProxy(fromInfo, passWd, valAcc.GetAddress().String(), "my memo", accInfo.GetAccountNumber(), sequence)
	require.NoError(t, err)
	fmt.Println(res)

}

func TestUnbindProxy(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Staking().UnbindProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)

}

func TestUnjail(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.129:21157", BroadcastBlock)
	client := NewClient(config)

	remoteValMnemo := "buzz solution music normal mom evolve online oxygen fox enhance atom fluid"
	fromInfo, _, err := utils.CreateAccountWithMnemo(remoteValMnemo, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Slashing().Unjail(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestSend(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Token().Send(fromInfo, passWd, targetAddr, "10.24okt", "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestMultiSend(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	transStr := `okchain1g7c3nvac7mjgn2m9mqllgat8wwd3aptdqket5k 1.024okt
okchain1aac2la53t933t265nhat9pexf9sde8kjnagh9m 2.048okt`
	transfers, err := utils.ParseTransfersStr(transStr)
	require.NoError(t, err)

	res, err := client.Token().MultiSend(fromInfo, passWd, transfers, "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestIssue(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:26657", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Token().Issue(fromInfo, passWd, "btc", "BitCoin", "100000000",
		"the token of Bitcoin", "my memo", true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)

}

func TestList(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:26657", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Dex().List(fromInfo, passWd, "btc-65f", "okt", "0.02", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestNewOrders(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:26657", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	//res, err := client.Order().NewOrders(
	//	fromInfo,
	//	passWd,
	//	"btc-9ec_okt,btc-9ec_okt,btc-9ec_okt",
	//	"BUY,SELL,BUY",
	//	"11.2,22.3,33.4",
	//	"1.23,2.34,3.45",
	//	"my memo",
	//	accInfo.GetAccountNumber(),
	//	accInfo.GetSequence())
	res, err := client.Order().NewOrders(
		fromInfo,
		passWd,
		"btc-65f_okt,btc-65f_okt,btc-65f_okt",
		"BUY,BUY,BUY",
		"11.2,22.3,33.4",
		"1.23,2.34,3.45",
		"my memo",
		accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
	fmt.Println("orderIds:", utils.GetOrderIDsFromResponse(&res))
}

func TestDeposit(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Dex().Deposit(fromInfo, passWd, "btc-216_okt", "1024.2048okt", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestWithdraw(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Dex().Withdraw(fromInfo, passWd, "btc-216_okt", "0.2048okt", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestGenerateUnsignedTransferOwnershipTx(t *testing.T) {
	config := NewClientConfig("", BroadcastBlock)
	client := NewClient(config)
	err := client.Dex().GenerateUnsignedTransferOwnershipTx("btc-216_okt", addr, targetAddr, "my memo", "./unsignedTx.json")
	require.NoError(t, err)
}

func TestMultiSign(t *testing.T) {
	config := NewClientConfig("", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	err = client.Dex().MultiSign(fromInfo, passWd, "./unsignedTx.json", "./signedTx.json")
	require.NoError(t, err)
}

func TestTransferOwnership(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:26657", BroadcastBlock)
	client := NewClient(config)
	// 1.generate unsigned transfer-ownership tx file
	err := client.Dex().GenerateUnsignedTransferOwnershipTx("btc-216_okt", addr, targetAddr, "my memo", "./unsignedTx.json")
	require.NoError(t, err)

	// 2.multi-sign the stdTx by the receiver
	recvInfo, _, err := utils.CreateAccountWithMnemo(targetMnemonic, name, passWd)
	require.NoError(t, err)
	err = client.Dex().MultiSign(recvInfo, passWd, "./unsignedTx.json", "./signedTx.json")
	require.NoError(t, err)

	// 3.transfer ownership with the signed tx file
	ownInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(ownInfo.GetAddress().String())
	require.NoError(t, err)

	res, err := client.Dex().TransferOwnership(ownInfo, passWd, "./signedTx.json", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

func TestCancelOrders(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := client.Auth().QueryAccount(fromInfo.GetAddress().String())
	require.NoError(t, err)

	orderIDs := "ID0000000055-1,ID0000000055-2,ID0000000055-3"
	res, err := client.Order().CancelOrders(fromInfo, passWd, orderIDs, "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}

// query test

func TestQueryValidators(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	vals, err := client.Staking().QueryValidators()
	require.NoError(t, err)
	for _, v := range vals {
		fmt.Println(v)
	}
}

func TestQueryValidator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	valAddr := "okchainvaloper1wsrrv0q4ldqjm2lxayuscwthcht55crdnt6her"
	val, err := client.Staking().QueryValidator(valAddr)
	require.NoError(t, err)
	fmt.Println(val)
}

func TestQueryDelegator(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	delResp, err := client.Staking().QueryDelegator(addr)
	require.NoError(t, err)
	fmt.Println(delResp)
}

func TestQueryTokenInfo(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	tokens, err := client.Token().QueryTokenInfo(addr, "")
	require.NoError(t, err)
	fmt.Println(tokens)

	tokens, err = client.Token().QueryTokenInfo("", "btc-32e")
	require.NoError(t, err)
	fmt.Println(tokens)

	tokens, err = client.Token().QueryTokenInfo(addr+"123", "btc-9ec")
	require.NoError(t, err)
	fmt.Println(tokens)

}

func TestQueryProducts(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.123:20157", BroadcastBlock)
	client := NewClient(config)
	tokenPairs, err := client.Dex().QueryProducts("", 1, 50)
	require.NoError(t, err)
	for _, v := range tokenPairs {
		fmt.Println(v)
	}
}

func TestQueryAccountTokensInfo(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	tokensInfo, err := client.Token().QueryAccountTokensInfo(addr)
	require.NoError(t, err)
	fmt.Println(tokensInfo)
}

func TestQueryAccountTokenInfo(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	tokensInfo, err := client.Token().QueryAccountTokenInfo(addr, "btc-e68")
	require.NoError(t, err)
	fmt.Println(tokensInfo)
}

func TestQueryDepthBook(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	tokensInfo, err := client.Order().QueryDepthBook("btc-216_okt")
	require.NoError(t, err)
	fmt.Println(tokensInfo)
}

// need test

func TestQueryCandles(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)
	candles, err := client.Backend().QueryCandles("tbtc-44f_tusdk-0cd", 60, 100)
	require.NoError(t, err)
	for _, line := range candles {
		fmt.Println(line)
	}
}

func TestQueryTickers(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)
	tickers, err := client.Backend().QueryTickers(10)
	require.NoError(t, err)
	for _, t := range tickers {
		fmt.Println(t)
	}
}

func TestQueryRecentTxRecord(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)
	record, err := client.Backend().QueryRecentTxRecord("tbtc-44f_tusdk-0cd", 0, int(time.Now().Unix()), 0, 10)
	require.NoError(t, err)
	for _, res := range record {
		fmt.Println(res)
	}
}

func TestQueryOpenOrders(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)

	product := "tbtc-44f_tusdk-0cd"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	openOrdersList, err := client.Backend().QueryOpenOrders(backendAddr, product, side, start, end, page, perPage)
	require.NoError(t, err)
	for _, order := range openOrdersList {
		fmt.Println(order)
	}
}

func TestQueryClosedOrders(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)

	product := "tbtc-44f_tusdk-0cd"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	closedOrdersList, err := client.Backend().QueryClosedOrders(backendAddr, product, side, start, end, page, perPage)
	require.NoError(t, err)
	for _, order := range closedOrdersList {
		fmt.Println(order)
	}
}

func TestQueryDeals(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)

	product := "tbtc-44f_tusdk-0cd"
	side := "BUY"
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	deals, err := client.Backend().QueryDeals(backendAddr, product, side, start, end, page, perPage)
	require.NoError(t, err)
	for _, deal := range deals {
		fmt.Println(deal)
	}
}

func TestQueryTransactions(t *testing.T) {
	config := NewClientConfig("tcp://192.168.13.125:20157", BroadcastBlock)
	client := NewClient(config)

	typeCode := 0
	start, end := 1, int(time.Now().Unix())
	page, perPage := 0, 10

	txs, err := client.Backend().QueryTransactions(backendAddr, typeCode, start, end, page, perPage)
	require.NoError(t, err)
	for _, tx := range txs {
		fmt.Println(tx)
	}
}
