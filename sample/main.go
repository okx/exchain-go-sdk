package main

import (
	sdk "github.com/okex/okchain-go-sdk"
	"github.com/okex/okchain-go-sdk/utils"
	"log"
)

const (
	// TODO: link to mainnet of OKChain later
	rpcURL = "tcp://127.0.0.1:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	// user's password
	passWd = "12345678"
	// target address
	addr     = "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"
	baseCoin = "okt"
)

func main() {
	//-------------------- 1. preparation --------------------//

	// create a client
	config, err := sdk.NewClientConfig(rpcURL, "okchain", sdk.BroadcastBlock, "0.01okt")
	if err != nil {
		log.Fatal(err)
	}

	cli := sdk.NewClient(config)

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

	//-------------------- 4. delegate for staking --------------------//

	// increase sequence number
	sequenceNum++
	res, err = cli.Staking().Delegate(fromInfo, passWd, "0.1"+baseCoin, "my memo", accountNum, sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

}
