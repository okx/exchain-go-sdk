package client

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"

	"testing"
)

const (
	name   = "alice"
	passWd = "12345678"
	// sender's mnemonic
	mnemonic = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	addr     = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	// target address
	mnemonic1 = "pepper basket run install fury scheme journey worry tumble toddler swap change"
	addr1     = "okchain1wux20ku36ntgtxpgm7my9863xy3fqs0xgh66d7"
	// validator address
	valAddr1 = "okchainvaloper10q0rk5qnyag7wfvvt7rtphlw589m7frs863s3m"
	valAddr2 = "okchainvaloper1g7znsf24w4jc3xfca88pq9kmlyjdare6mph5rx"
	valAddr3 = "okchainvaloper1alq9na49n9yycysh889rl90g9nhe58lcs50wu5"
	valAddr4 = "okchainvaloper1svzxp4ts5le2s4zugx34ajt6shz2hg42a3gl7g"
	// validator mnemonic
	valMnemonic = "race imitate stay curtain puppy suggest spend toe old bridge sunset pride"
	valName     = "validator"
)

func TestOKChainClient_Send(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Send(fromInfo, passWd, addr1, "10.24okt", "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_NewOrders(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.NewOrders(
		fromInfo,
		passWd,
		"xxb-031_okt,xxb-031_okt,xxb-031_okt",
		"BUY,SELL,BUY",
		"11.2,22.3,33.4",
		"1.23,2.34,3.45",
		"my memo",
		accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
	fmt.Println("orderIds:", getOrderIdsFromResponse(&res))
}

func TestOKChainClient_CancelOrders(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	orderIds := "ID0000003438-1,ID0000003438-3"
	res, err := cli.CancelOrders(fromInfo, passWd, orderIds, "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_Delegate(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Delegate(fromInfo, passWd, "1024.2048okt", "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_Unbond(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Unbond(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_Vote(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	// delegate first
	sequence := accInfo.GetSequence()
	_, err = cli.Delegate(fromInfo, passWd, "1024.2048okt", "my memo", accInfo.GetAccountNumber(),
		sequence)
	assertNotEqual(t, err, nil)

	// vote then
	sequence++
	valsToVoted := []string{valAddr1, valAddr2, valAddr3, valAddr4}
	res, err := cli.Vote(fromInfo, passWd, valsToVoted, "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_DestroyValidator(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(valMnemonic, valName, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.DestroyValidator(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_Unjail(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(valMnemonic, valName, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Unjail(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_MultiSend(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	transStr := `okchain1g7c3nvac7mjgn2m9mqllgat8wwd3aptdqket5k 1.024okt
okchain1aac2la53t933t265nhat9pexf9sde8kjnagh9m 2.048okt`
	transfers, err := utils.ParseTransfersStr(transStr)
	assertNotEqual(t, err, nil)

	res, err := cli.MultiSend(fromInfo, passWd, transfers, "my memo", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_CreateValidator(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	pubkeyStr := "okchainvalconspub1zcjduepqghrtvkngejwese62wg49ewskz4r93vkyj3md5mg5rf7twcc6jduqpqw66q"
	res, err := cli.CreateValidator(fromInfo, passWd, pubkeyStr, "my moniker", "my identity",
		"my website", "my details", "1okt", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_EditValidator(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(valMnemonic, valName, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.EditValidator(fromInfo, passWd, "my moniker", "my identity", "my website",
		"my details", "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_RegisterProxy(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	sequence := accInfo.GetSequence()
	res, err := cli.Delegate(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)

	sequence++
	res, err = cli.RegisterProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_UnregisterProxy(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.UnregisterProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_BindProxy(t *testing.T) {
	cli := NewClient(rpcUrl)

	// validator becomes a proxy
	valAcc, _, err := utils.CreateAccountWithMnemo(valMnemonic, valName, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(valAcc.GetAddress().String())
	assertNotEqual(t, err, nil)

	sequence := accInfo.GetSequence()
	res, err := cli.Delegate(valAcc, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)

	sequence++
	res, err = cli.RegisterProxy(valAcc, passWd, "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)

	// delegator tries to bind proxy
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err = cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	sequence = accInfo.GetSequence()
	res, err = cli.Delegate(fromInfo, passWd, "102.4okt", "my memo", accInfo.GetAccountNumber(), sequence)
	assertNotEqual(t, err, nil)

	sequence++
	res, err = cli.BindProxy(fromInfo, passWd, valAcc.GetAddress().String(), "my memo", accInfo.GetAccountNumber(),
		sequence)
	assertNotEqual(t, err, nil)
	fmt.Println(res)

}

func TestOKChainClient_UnbindProxy(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.UnbindProxy(fromInfo, passWd, "my memo", accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)

}

func TestOKChainClient_Issue(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Issue(fromInfo, passWd, "btc", "BitCoin", "100000000",
		"the token of Bitcoin", "my memo", true, accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)

}

func TestOKChainClient_List(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.List(fromInfo, passWd, "btc-e68", "okt", "0.02", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_Deposit(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Deposit(fromInfo, passWd, "btc-e68_okt", "1024.2048okt", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_Withdraw(t *testing.T) {
	cli := NewClient(rpcUrl)
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	assertNotEqual(t, err, nil)
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	assertNotEqual(t, err, nil)

	res, err := cli.Withdraw(fromInfo, passWd, "btc-e68_okt", "0.2408okt", "my memo",
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	assertNotEqual(t, err, nil)
	fmt.Println(res)
}

func TestOKChainClient_TransferOwnership(t *testing.T) {
	// 1.generate unsigned transfer-ownership tx file
	err := utils.GenerateUnsignedTransferOwnershipTx("btc-e68_okt", addr, addr1, "my memo", "./unsignedTx.json")
	require.NoError(t, err)

	// 2.multi-sign the stdTx by the receiver
	recvInfo, _, err := utils.CreateAccountWithMnemo(mnemonic1, name, passWd)
	require.NoError(t, err)
	err = utils.MultiSign(recvInfo, passWd, "./unsignedTx.json", "./signedTx.json")
	require.NoError(t, err)

	// 3.transfer ownership with the signed tx file
	cli := NewClient(rpcUrl)
	ownInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)
	accInfo, err := cli.GetAccountInfoByAddr(ownInfo.GetAddress().String())
	require.NoError(t, err)
	res, err := cli.TransferOwnership(ownInfo, passWd, "./signedTx.json", accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	fmt.Println(res)
}
