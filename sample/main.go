package main

import (
	"fmt"
	"github.com/okex/okchain-go-sdk/client"
	"github.com/okex/okchain-go-sdk/utils"
	"log"
)

const (
	rpcUrl = "3.13.150.20:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = "total lottery arena when pudding best candy until army spoil drill pool"
	// user's password
	passWd = "12345678"
	// target address
	addr = "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"
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
	// sequence number of the account must be increased by 1 whenever a transaction of the account takes effect
	sequenceNum := accInfo.GetSequence()
	res, err := cli.Send(fromInfo, passWd, addr, "1tokt", "my memo", accInfo.GetAccountNumber(), sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	/* 4. place an order on OK Dex */
	sequenceNum++
	res, err = cli.NewOrder(fromInfo, passWd, "tokt_tusdk", "BUY", "1", "1", "my memo", accInfo.GetAccountNumber(), sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	orderId := res.Tags[1].Value
	fmt.Println(res)

	fmt.Println("orderId:", orderId)

	/* 5. cancel the order on OK Dex by orderID */
	sequenceNum++
	res, err = cli.CancelOrder(fromInfo, passWd, orderId, "my memo", accInfo.GetAccountNumber(), sequenceNum)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

	return
}
