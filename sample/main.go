package main

import (
	"log"

	"github.com/okex/okexchain-go-sdk"
	"github.com/okex/okexchain-go-sdk/utils"
)

const (
	// TODO: link to mainnet of OKExChain later
	rpcURL = "tcp://127.0.0.1:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	// user's password
	passWd = "12345678"
	// target address
	addr     = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	baseCoin = "okt"
)

func main() {
	//-------------------- 1. preparation --------------------//
	// NOTE: either of the both ways below to pay fees is available

	// WAY 1: create a client config with fixed fees
	config, err := gosdk.NewClientConfig(rpcURL, "okexchain", gosdk.BroadcastBlock, "0.01okt", 200000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}

	// WAY 2: alternative client config with the fees by auto gas calculation
	config, err = gosdk.NewClientConfig(rpcURL, "okexchain", gosdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	// create an account with your own mnemonic，name and password
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	if err != nil {
		log.Fatal(err)
	}

	//-------------------- 2. query for the information of your address --------------------//

	accInfo, err := cli.Auth().QueryAccount(fromInfo.GetAddress().String())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(accInfo)

	//-------------------- 3. transfer to other address --------------------//

	// sequence number of the account must be increased by 1 whenever a transaction of the account takes effect
	accountNum, sequenceNum := accInfo.GetAccountNumber(), accInfo.GetSequence()
	res, err := cli.Token().Send(fromInfo, passWd, addr, "1"+baseCoin, "my memo", accountNum, sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

	//-------------------- 4. deposit for staking --------------------//

	// increase sequence number
	sequenceNum++
	res, err = cli.Staking().Deposit(fromInfo, passWd, "0.1"+baseCoin, "my memo", accountNum, sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
