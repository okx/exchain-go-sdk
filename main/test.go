package main

import (
	bytes2 "bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"math/rand"
	"sync"
	"time"
)



/*
import (
	"fmt"
	sdk "github.com/okex/exchain-go-sdk"
)


func main() {
	rpcURL   := "http://127.0.0.1:8545"
	toAddr     := "0x79e1ae352245F72BDceB0cCBc2eF414bC07F154c"
	nonce	 :=	5
	privHex := "89c81c304704e9890025a5a91898802294658d6e4034a11c6116f4b129ea12d3"

	// build the client with own config
	config, _ := sdk.NewClientConfig(rpcURL, "exchainevm-8", sdk.BroadcastBlock, "0.01okt", 20000, 0, "")
	client := sdk.NewClient(config)

	res, err := client.Evm().SendTxEthereum(privHex, toAddr, "", "", 21000, uint64(nonce))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
*/

var (
	host    = "http://localhost:8545"
//	host	= "http://18.167.77.79:26659"
	privKey = "e47a1fe74a7f9bfa44a362a3c6fbe96667242f62e6b8e138b3f61bd431c3215d"
	toPrivKey = "89c81c304704e9890025a5a91898802294658d6e4034a11c6116f4b129ea12d3"

	sampleContractByteCode []byte
	sampleContractABI      abi.ABI
)

func init() {
	bin, err := ioutil.ReadFile("/Users/oker/go/src/github.com/okex/exchain-web3-sample/sample_contract/Storage.bin")
	if err != nil {
		fmt.Println(err)
	}
	sampleContractByteCode = common.Hex2Bytes(string(bin))

	abiByte, err := ioutil.ReadFile("/Users/oker/go/src/github.com/okex/exchain-web3-sample/sample_contract/Storage.abi")
	if err != nil {
		fmt.Println(err)
	}
	sampleContractABI, err = abi.JSON(bytes2.NewReader(abiByte))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	client, _ := ethclient.Dial(host)

	chainID, _ := client.NetworkID(context.Background())
	privateKey, _ := crypto.HexToECDSA(privKey)
	gasPrice, _ := client.SuggestGasPrice(context.Background())

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		fmt.Println("failed to switch unencrypted private key -> secp256k1 private key: %+v", err)
	}
	// 0.4 secp256k1 private key -> pubkey -> address
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
	/*
	unsignedTx := types.NewTransaction(6, common.HexToAddress("0x2CF4ea7dF75b513509d95946B43062E26bD88035"), big.NewInt(1000000000000000000), 30000, gasPrice, []byte{})
	signedTx, _ := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), privateKey)
	_ = client.SendTransaction(context.Background(), signedTx)

	tx := new(evmtypes.MsgEthereumTx)
	// RLP decode raw transaction bytes
	data, _ := rlp.EncodeToBytes(tx)
	_ = rlp.DecodeBytes(data, tx)
	txEncoder := authclient.GetTxEncoder(nil)
	txBytes, _ := txEncoder(tx)
	txHash := common.HexToHash(strings.ToUpper(hex.EncodeToString(tmhash.Sum(txBytes))))
	fmt.Println(txHash)
	*/
}
/*
func mainbak() {
	//
	// 0. init
	//
	// 0.1 init client
	client, err := ethclient.Dial(host)
	if err != nil {
		fmt.Println("failed to initialize client: %+v", err)
	}
	// 0.2 get the chain-id from network
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("failed to fetch the chain-id from network: %+v", err)
	}
	// 0.3 unencrypted private key -> secp256k1 private key
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		fmt.Println("failed to switch unencrypted private key -> secp256k1 private key: %+v", err)
	}
	// 0.4 secp256k1 private key -> pubkey -> address
	pubkey := privateKey.Public()
	pubkeyECDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("failed to switch secp256k1 private key -> pubkey")
	}
	fromAddress := crypto.PubkeyToAddress(*pubkeyECDSA)

	// 0.5 get the gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	//
	// 1. deploy contract
	//
	contractAddr := deployContract(client, fromAddress, gasPrice, chainID, privateKey)
	//
	// 2. call contract(write)
	//
	writeContract(client, fromAddress, gasPrice, chainID, privateKey, contractAddr)
	time.Sleep(time.Second * 3)
	//
	// 3. call contract(read)
	//
	readContract(client, contractAddr)
}

func deployContract(client *ethclient.Client,
	fromAddress common.Address,
	gasPrice *big.Int,
	chainID *big.Int,
	privateKey *ecdsa.PrivateKey) (contractAddr common.Address) {
	// 0. get the value of nonce, based on address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("failed to fetch the value of nonce from network: %+v", err)
	}

//	nonce++
	//1. simulate unsignedTx as you want, fill out the parameters into a unsignedTx
	unsignedTx := deployContractTx(nonce, gasPrice)

	// 2. sign unsignedTx -> rawTx
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("failed to sign the unsignedTx offline: %+v", err)
	}

	// 3. send rawTx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}

	// 4. get the contract address based on tx hash
	hash := getTxHash(signedTx)
	time.Sleep(time.Second * 5)

	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		fmt.Println(err)
	}

	return receipt.ContractAddress
}

func deployContractTx(nonce uint64, gasPrice *big.Int) *types.Transaction {
	value := big.NewInt(0)
	gasLimit := uint64(3000000)

	// Constructor
	input, err := sampleContractABI.Pack("")
	if err != nil {
		fmt.Println(err)
	}
	data := append(sampleContractByteCode, input...)
	return types.NewContractCreation(nonce, value, gasLimit, gasPrice, data)
}

func writeContract(client *ethclient.Client,
	fromAddress common.Address,
	gasPrice *big.Int,
	chainID *big.Int,
	privateKey *ecdsa.PrivateKey,
	contractAddr common.Address) {
	// 0. get the value of nonce, based on address
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("failed to fetch the value of nonce from network: %+v", err)
	}

	unsignedTx := writeContractTx(nonce, contractAddr, gasPrice)
	// 2. sign unsignedTx -> rawTx
	signedTx, err := types.SignTx(unsignedTx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("failed to sign the unsignedTx offline: %+v", err)
	}

	// 3. send rawTx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
	}
}

func writeContractTx(nonce uint64, contractAddr common.Address, gasPrice *big.Int) *types.Transaction {
	value := big.NewInt(0)
	gasLimit := uint64(3000000)

	num := big.NewInt(999)
	data, err := sampleContractABI.Pack("store", num)
	if err != nil {
		fmt.Println(err)
	}
	return types.NewTransaction(nonce, contractAddr, value, gasLimit, gasPrice, data)
}

func readContract(client *ethclient.Client, contractAddr common.Address) {
	data, err := sampleContractABI.Pack("retrieve")
	if err != nil {
		fmt.Println(err)
	}

	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	output, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		panic(err)
	}

	ret, err := sampleContractABI.Unpack("retrieve", output)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}

func getTxHash(signedTx *types.Transaction) common.Hash {
	ts := types.Transactions{signedTx}
	rawTx := hex.EncodeToString(ts.GetRlp(0))

	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		fmt.Println(err)
	}

	tx := new(evmtypes.MsgEthereumTx)
	// RLP decode raw transaction bytes
	if err := rlp.DecodeBytes(rawTxBytes, tx); err != nil {
		fmt.Println(err)
	}

	cdc := codec.MakeCodec(app.ModuleBasics)
	txEncoder := authclient.GetTxEncoder(cdc)
	txBytes, err := txEncoder(tx)
	if err != nil {
		fmt.Println(err)
	}

	var hexBytes bytes.HexBytes
	hexBytes = tmhash.Sum(txBytes)
	hash := common.HexToHash(hexBytes.String())
	return hash
}
*/