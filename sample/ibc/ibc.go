package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	gosdk "github.com/okex/exchain-go-sdk"
	"github.com/okex/exchain-go-sdk/utils"
	secp256k12 "github.com/okex/exchain/libs/cosmos-sdk/crypto/keys/ibc-key"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"log"
)

const (
	rpcURL = "tcp://127.0.0.1:36657"
	// user's name
	name = "admin17"
	// user's mnemonic
	mnemonic = "antique onion adult slot sad dizzy sure among cement demise submit scare"
	// user's password
	passWd = "12345678"
	// target address
	addr            = "cosmos1n064mg7jcxt2axur29mmek5ys7ghta4u4mhcjp"
	baseCoin        = "okt"
	aliceKey string = "e47a1fe74a7f9bfa44a362a3c6fbe96667242f62e6b8e138b3f61bd431c3215d"
)

func main() {
	//-------------------- 1. preparation --------------------//
	// NOTE: either of the both ways below to pay fees is available

	// WAY 1: create a client config with fixed fees
	//config, err := gosdk.NewClientConfig(rpcURL, "exchain-101", gosdk.BroadcastSync, "0.000001okt", 10000,
	//	0, "")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// WAY 2: alternative client config with the fees by auto gas calculation
	config, err := gosdk.NewClientConfig(rpcURL, "exchain-101", gosdk.BroadcastSync, "0.000001okt", 450000,
		1.1, "0.000000000000000012okt")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	// create an account with your own mnemonic，name and password

	if err != nil {
		log.Fatal(err)
	}

	//-------------------- 2. query for the information of your address --------------------//

	accInfo, err := cli.Auth().QueryAccount("ex1hr26cyc335g7p5e948a7vkmwnx3fmxfzwdyryf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(accInfo)

	//-------------------- 3. ibc transfer to the address --------------------//

	// sequence number of the account must be increased by 1 whenever a transaction of the account takes effect
	//accountNum, sequenceNum := accInfo.GetAccountNumber(), accInfo.GetSequence()
	priStr, err := utils.GeneratePrivateKeyFromMnemo(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	key, err := crypto.HexToECDSA(priStr)
	if err != nil {
		log.Fatal(err)
	}

	d := crypto.FromECDSA(key)

	fee := sdk.NewCoinAdapter("okt", sdk.NewInt(1))
	fees := []sdk.CoinAdapter{fee}

	res, err := cli.Ibc().Transfer(secp256k12.GenPrivKeyFromSecret(d), "channel-0", addr, "1000okt", fees, "memo", "http://127.0.0.1:16657")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

}
