## OKChain Go SDK

[![CircleCI](https://circleci.com/gh/okex/okchain-go-sdk/tree/master.svg?style=shield)](https://circleci.com/gh/okex/okchain-go-sdk/tree/master)
[![codecov](https://codecov.io/gh/okex/okchain-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/okex/okchain-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/okex/okchain-go-sdk)](https://goreportcard.com/report/github.com/okex/okchain-go-sdk)
[![license](https://img.shields.io/github/license/okex/okchain-go-sdk.svg)](https://github.com/okex/okchain-go-sdk/blob/master/LICENSE)
[![LoC](https://tokei.rs/b1/github/okex/okchain-go-sdk)](https://github.com/okex/okchain-go-sdk)
[![GolangCI](https://golangci.com/badges/github.com/okex/okchain-go-sdk.svg)](https://golangci.com/r/github.com/okex/okchain-go-sdk)

The OKChain Go SDK is a lightweight Go library to interact with OKChain.

### 1. Components

- client.go - The main client of gosdk is created in this file. Users is supposed to set up the config with own requirement during the client creation.
- expose - Abstraction with the interface of each module. The implements of it are filled in the folder `module`.

- module - The main logic for GO SDK queries and txs are classfied by their own module name in okchain. Developers can find out the concrete design under the specific module folder. Please focus on the files, tx.go and query.go. 
- mocks - Mock client tools for unit test of the main client in GO SDK.
- sample - A clear short user guild is showed here.
-  types - The necessary struct set of okchain is built here. Developers are allowed to import some basic types like Dec and AccAddress directly if they want.
- utils -  A useful tool set for the one who is going to send more transcations and queries is spilted by module name as the file name. Beyond that, the operation of account keys with mnemonics remains in file `account.go`.

### 2. Installation

Go version above 1.12 is required.

The developer can get the OKChain Go SDK directly by `git clone` from github : https://github.com/okex/okchain-go-sdk

### 3. API

The api functions of transactions and queries are all under the path `expose`. You can find more details in okchain-docs : https://okchain-docs.readthedocs.io/zh_CN/latest/api/sdk/go-sdk.html

### 4. Tendermint query

OKChain Go SDK also provides node query functions so that the underlying information of the blockchain is available for developers.

The tendermint query functions could be found in the file `exposed/tendermint.go `. Developers could make it through with the file `module/tendermint/query.go` and get clear that how to invoke them.

### 5. Example

`Client` seems necessary to every operation with Go SDK. Here are the examples :

```go
// rpcURL should modified according to the actual situation
rpcURL	 := "3.13.150.20:26657"
name     := "alice"
passWd   := "12345678"
mnemonic := "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
addr1    := "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"

// build the client with own config
	config, err := sdk.NewClientConfig(rpcURL, "okchain", BroadcastBlock, "0.01okt", 20000)
	require.NoError(t, err)
	client := sdk.NewClient(config)

	// create your account key info by 'name','passWd' and 'mnemonic'
	keyInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	// get info of your account from OKChain
	accInfo, err := client.Auth().QueryAccount(keyInfo.GetAddress().String())
	require.NoError(t, err)

	// transfer some okt to addr1
	res, err := client.Token().Send(keyInfo, passWd, addr1, "0.1024okt", "my memno", accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
```

You can invoke more and more api functions with the object `client`.

### 6. Testing

All changes and addition of codes will be pushed with unit tests strictly. 

### 7. Contributing

No doubt that it's admirable to make contributions to OKChain Go SDK. You can provide your code as long as you have tested it with a local client and your unit test showed its validity.  

