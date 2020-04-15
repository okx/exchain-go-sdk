package sdk

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

// just a temporary test

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

func TestQueryValidators(t *testing.T) {
	config := NewClientConfig("tcp://127.0.0.1:10057", BroadcastBlock)
	client := NewClient(config)
	vals, err := client.Staking().QueryValidators()
	require.NoError(t, err)
	for _, v := range vals {
		fmt.Println(v)
	}
}

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
