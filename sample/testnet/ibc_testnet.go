package main

import (
	"fmt"
	gosdk "github.com/okex/exchain-go-sdk"
	"github.com/okex/exchain-go-sdk/utils"
	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	client_types "github.com/okx/okbchain/libs/ibc-go/modules/core/02-client/types"

	"log"
)

const (
	rpcURL = "http://18.177.169.67:26657"
	// user's name
	name = "admin17"
	// user's mnemonic
	mnemonic = "antique onion adult slot sad dizzy sure among cement demise submit scare"
	// user's password
	passWd = "12345678"
	// target address
	addr            = "cosmos1hkxc5nkcqyu7g9z2efkxq5gvy0trjdcr7fkzen"
	baseCoin        = "okt"
	aliceKey string = "e47a1fe74a7f9bfa44a362a3c6fbe96667242f62e6b8e138b3f61bd431c3215d"
)

func main() {

	config, err := gosdk.NewClientConfig(rpcURL, "exchain-65", gosdk.BroadcastSync, "0.000001okt", 450000,
		1.1, "0.000000000000000012okt")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	// create an account with your own mnemonic，name and password

	if err != nil {
		log.Fatal(err)
	}

	// sequence number of the account must be increased by 1 whenever a transaction of the account takes effect
	privateKey, err := utils.GenerateEthPrivateKeyFromMnemo(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	fee := sdk.NewCoinAdapter("wei", sdk.NewInt(45000000000000))
	fees := []sdk.CoinAdapter{fee}

	fmt.Println(privateKey.PubKey().Address())

	res, err := cli.Ibc().Transfer(privateKey, "channel-2", addr, "0.1okt", fees, "memo", client_types.Height{RevisionNumber: 1, RevisionHeight: 1000000})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
