package main

import (
	"fmt"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"log"

	gosdk "github.com/okex/exchain-go-sdk"
)

const (
	// TODO: link to mainnet of ExChain later
	rpcURL = "tcp://54.238.76.5:26657"
	// user's name
	name = "alice"
	// user's mnemonic
	mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
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
	config, err := gosdk.NewClientConfig(rpcURL, "exchain-65", gosdk.BroadcastBlock, "0.01okt", 200000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}

	// WAY 2: alternative client config with the fees by auto gas calculation
	config, err = gosdk.NewClientConfig(rpcURL, "exchain-65", gosdk.BroadcastBlock, "", 200000,
		1.1, "0.000000001okt")
	if err != nil {
		log.Fatal(err)
	}

	cli := gosdk.NewClient(config)

	fmt.Println("=================================================QueryFeesplits========================================================")
	queryFeesplits, err := cli.Feesplit().QueryFeesplits(&query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryFeesplits)

	fmt.Println("=================================================QueryFeeSplit========================================================")
	queryFeesplit, err := cli.Feesplit().QueryFeeSplit("0x0554c61F21936dAD6b1F5bDc685e266beBd04234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryFeesplit)

	fmt.Println("=================================================QueryDeployerFeeSplits========================================================")
	queryDeployerFeeSplits, err := cli.Feesplit().QueryDeployerFeeSplits("ex1l5jugfjaqys4k64rpqud3lymf8a3csg6ds2j4h", &query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryDeployerFeeSplits)

	fmt.Println("=================================================QueryWithdrawerFeeSplits========================================================")
	queryWithdrawerFeeSplits, err := cli.Feesplit().QueryWithdrawerFeeSplits("ex1xezsnqfln9vcujhaalwrvl96ukyq9q2gfh8gek", &query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      100,
		CountTotal: false,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryWithdrawerFeeSplits)

	fmt.Println("=================================================QueryParams========================================================")
	queryParams, err := cli.Feesplit().QueryParams()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(queryParams)

}
