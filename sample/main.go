package main

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/client"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/prometheus/common/log"
)

const (
	rpcUrl   = "3.13.150.20:26657"
	name     = "alice"
	mnemonic = "total lottery arena when pudding best candy until army spoil drill pool"
	passWd   = "12345678"
	addr     = "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"
)

func main() {
	/* 1. preparation */
	// create a client
	cli := client.NewClient(rpcUrl)
	// create an account with your own mnemonic，name and password
	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	if err != nil {
		log.Fatal(err)
	}

	/* 2. query for the information of your address */
	accInfo, err := cli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(accInfo)

	/* 3. transfer to other address */
	res, err := cli.Send(fromInfo, passWd, addr, "1tokt", "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	/* 4. place an order on OK Dex */
	res, err = cli.NewOrder(fromInfo, passWd, "tokt_tusdk", "BUY", "1", "1", "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
	if err != nil {
		log.Fatal(err)
	}

	orderId := res.Tags[1].Value
	fmt.Println(res)

	fmt.Println("orderId:", orderId)

	/* 5. cancel the order on OK Dex by orderID */
	res, err = cli.CancelOrder(fromInfo, passWd, orderId, "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return
}
