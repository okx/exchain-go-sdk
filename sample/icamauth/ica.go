package main

import (
	gosdk "github.com/okex/exchain-go-sdk"
	"github.com/okex/exchain-go-sdk/utils"

	"log"
)

const (
	rpcURL = "https://exchaintesttmrpc.okex.org"
	// user's name
	name = "admin17"
	// user's mnemonic
	mnemonic = "antique onion adult slot sad dizzy sure among cement demise submit scare"
	// user's password
	passWd = "12345678"
	// target address
	addr                = "cosmos1n064mg7jcxt2axur29mmek5ys7ghta4u4mhcjp"
	baseCoin            = "okt"
	privateKey   string = "8ff3ca2d9985c3a52b459e2f6e7822b23e1af845961e22128d5f372fb9aa5f17"
	connectionID        = "connection-73"
)

func main() {

	config, err := gosdk.NewClientConfig(rpcURL, "exchain-65", gosdk.BroadcastSync, "0.000001okt", 450000,
		1.1, "0.000000000000000012okt")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	// create an account with your own privateKey name and password
	fromInfo, err := utils.CreateAccountWithPrivateKey(privateKey, "admin", passWd)
	if err != nil {
		log.Fatal(err)
	}

	accInfo, err := cli.Auth().QueryAccount(fromInfo.GetAddress().String())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(accInfo)

	str := `{
  "@type":"/cosmos.bank.v1beta1.MsgSend",
  "from_address":"cosmos1km7ylyepj9h4afusjcx472q4ml9l9s8mzhmhymvxmvtn2hqxq32s53nlmg",
  "to_address":"cosmos10h9stc5v6ntgeygf5xf945njqq5h32r53uquvw",
  "amount": [
    {
      "denom": "uatom",
      "amount": "1000"
    }
  ]
}`
	response, err := cli.Icamauth().SubmitTx(fromInfo, passWd, connectionID, "", []byte(str), accInfo.GetAccountNumber(), accInfo.GetSequence())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
