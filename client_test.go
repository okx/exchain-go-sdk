package sdk

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

// just a temporary test, it will be removed later

const (
	name   = "alice"
	passWd = "12345678"
	// sender's mnemonic
	mnemonic   = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	addr       = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	targetAddr = "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"
)

// transact tx

func TestDelegate(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
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
