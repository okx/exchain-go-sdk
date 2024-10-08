package main

import (
	"log"

	gosdk "github.com/okex/exchain-go-sdk"
	"github.com/okex/exchain-go-sdk/utils"
)

const (
	// TODO: link to mainnet of ExChain later
	rpcURL = "tcp://127.0.0.1:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = ""
	// user's password
	passWd = "12345678"
	// target address
	addr     = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	baseCoin = "okt"
)

func main() {
	//-------------------- 1. preparation --------------------//
	// NOTE: either of the both ways below to pay fees is available

	// WAY 1: create a client config with fixed fees
	config, err := gosdk.NewClientConfig(rpcURL, "exchain", gosdk.BroadcastBlock, "0.01okt", 200000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}

	// WAY 2: alternative client config with the fees by auto gas calculation
	config, err = gosdk.NewClientConfig(rpcURL, "exchain", gosdk.BroadcastBlock, "", 200000,
		1.1, "0.000000001okt")
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
