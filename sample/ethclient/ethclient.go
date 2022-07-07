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
	alice    string = "0xaD37A476c7D3b8F382C5DfC5E789e6540ea246bb"
	bob      string = "0x0073F2E28ef8F117e53d858094086Defaf1837D5"
	aliceKey string = "5a72d444804664c3cf38fffc6117e6142146ddac25abaa35b72eb86dfe6ae56c"
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

	blockNumber, err := client.BlockNumber(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blockNumber:", blockNumber)

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(alice), big.NewInt(int64(blockNumber)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance:", balance)

	nonce, err := client.NonceAt(context.Background(), common.HexToAddress(alice), big.NewInt(int64(blockNumber)))
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
	code, err := client.CodeAt(context.Background(), common.HexToAddress(alice), big.NewInt(int64(blockNumber)))
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

	re, err := client.CallContract(context.Background(), msg, big.NewInt(int64(blockNumber)))
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
