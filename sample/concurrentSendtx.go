package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

func main1() {
	client, _ := ethclient.Dial(host)

	chainID, _ := client.NetworkID(context.Background())
	privateKey, _ := crypto.HexToECDSA(aliceKey)
	gasPrice, _ := client.SuggestGasPrice(context.Background())

	privateKey, err := crypto.HexToECDSA(aliceKey)
	if err != nil {
		fmt.Println("failed to switch unencrypted private key -> secp256k1 private key: %+v", err)
	}

	pubkey := privateKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("failed to switch secp256k1 private key -> pubkey")
	}
	toAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	a := make([]int, 3000)
	list := rand.Perm(len(a))
	var wg sync.WaitGroup
	wg.Add(len(list))
	fmt.Println(list)
	for i := 0; i < len(list); i++ {
		n := list[i] + 0
		go func(nonce uint64) {
			defer wg.Done()
			unsignedTx := types.NewTransaction(nonce, toAddress, big.NewInt(1000000000000000000), 30000, gasPrice, []byte{})
			signedTx, _ := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), privateKey)
			err := client.SendTransaction(context.Background(), signedTx)
			for err != nil {
				fmt.Println(err)
				time.Sleep(time.Millisecond * 100)
				err = client.SendTransaction(context.Background(), signedTx)
			}
		}(uint64(n))
	}
	wg.Wait()
}