package main

import (
	"context"
	"fmt"
	"log"

	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	gosdk "github.com/okex/exchain-go-sdk"
)

const (
	host     string = "https://exchaintestrpc.okex.org"
	alice    string = "0x2CF4ea7dF75b513509d95946B43062E26bD88035"
	bob      string = "0x0073F2E28ef8F117e53d858094086Defaf1837D5"
	aliceKey string = "e47a1fe74a7f9bfa44a362a3c6fbe96667242f62e6b8e138b3f61bd431c3215d"
)

func main() {

	client, err := gosdk.NewEthClient(context.Background(), host)
	if err != nil {
		fmt.Println(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice", gasPrice)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(alice), big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance:", balance)

	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(alice), big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nonce", nonce)

	pendingNonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(alice))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pendingNonce", pendingNonce)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainID", chainID)

	privateKey, _ := crypto.HexToECDSA(aliceKey)
	unsignedTx := types.NewTransaction(pendingNonce, common.HexToAddress(bob), big.NewInt(1000000000000000000), 30000, gasPrice, []byte{})
	signedTx, _ := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), privateKey)

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 4)
	fmt.Println("txHash", signedTx.Hash())

	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("recipt %+v\n", receipt)

	pendingCode, err := client.PendingCodeAt(context.Background(), common.HexToAddress(alice))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("pendingCode", pendingCode)
	code, err := client.CodeAt(context.Background(), common.HexToAddress(alice), big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("code", code)

	to := common.HexToAddress(bob)
	msg := ethereum.CallMsg{From: common.HexToAddress(alice), To: &to, GasPrice: gasPrice, Value: big.NewInt(1), Data: []byte{}}
	estimateGas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("estimateGas", estimateGas)

	re, err := client.CallContract(context.Background(), msg, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("callContract", re)

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block", block)

	ss, err := client.SubscribeNewHead(context.Background(), make(chan *types.Header))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("subscription", ss)
	}
	es, err := client.EthSubscribe(context.Background(), make(chan *types.Header), "newHeads")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ethSubscribe", es)
	}

	var hex hexutil.Big
	err = client.CallContext(context.Background(), &hex, "eth_gasPrice")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CallContext", hex)
}
