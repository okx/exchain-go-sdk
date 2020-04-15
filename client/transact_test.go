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







func TestOKChainClient_TransferOwnership(t *testing.T) {
	// 1.generate unsigned transfer-ownership tx file
	err := utils.GenerateUnsignedTransferOwnershipTx("btc-216_okt", addr, addr1, "my memo", "./unsignedTx.json")
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
