## OKChain Go SDK

The OKChain Go SDK is a lightweight Go library to interact with OKChain.

### 1. Components

- okclient - The functions which could be invoked to query information and transact trades are designed as class methods of `okclient` inside. An object of the struct `okclient` should be created first before the voking of api functions.
- common -  It contains some third-party libs and definitions of parameters which are used in the api functions.
- crypto - The path of cryptography, encoding and keys. The creation of accounts is relied on it closely.
- types - All the definitions of interfaces and structs which work during the  interaction between OKChain Go SDK and OKChain are here inside. The change will cause the failing interaction. So it is not recommended to modify any code under this path.
- utils - some tool functions. The developer should focus on it if they want to be clear about the operations of ok accounts.
- vendor - Third-party dependency libs such as the rpc client core.  

There are some test modules in path `okclient` and `utils` as well. The developer will know how to design the code themselves by checking the test code and running the test modules.

- okclient/query.go : all the api functions of querying information from OKChain, such as getting infomation of tokens and K lines,  are here.
- okclient/transact.go : all the api functions of transact a trade to OKChain,  such as transfering and pending a new order , are here. It is noticed that the fee is required no matter what you do.

### 2. Installation

Go version above 1.12 is required.

The developer can install the OKChain Go SDK by `git clone` from github : https://github.com/okex/okchain-go-sdk

### 3. API

The api functions of querying and transacting are in the files 'okclient/query,go'  and `okclient/transact.go`. You can find more details in okchain-docs : https://okchain-docs.readthedocs.io/zh_CN/latest/api/sdk/go-sdk.html

### 4. Node query

 The OKChain Go SDK also provides node query functions so that the underlying information of the blockchain can be got by developers.

The node query functions could be found in file `okclient/node_query.go `. The module of them are in file `okclient/node_query_test.go`. The develop could make it through and get clear that how to invoke them.

More details of node querying from OKChain in okchain-docs : https://okchain-docs.readthedocs.io/zh_CN/latest/api/node_rpc.html

### 5. Example

Every operation by using Go SDK needs `okclient`. Here are the examples :

```go
// rpcUrl can be modified according to the actual situation
rpcUrl	 := "3.13.150.20:26657"
okCli 	 := NewClient(rpcUrl)
name     := "alice"
passWd   := "12345678"
mnemonic := "sustain hole urban away boy core lazy brick wait drive tiger tell"
addr1    := "okchain1hw4r48aww06ldrfeuq2v438ujnl6alszzzqpph"

// create your account key info by 'name','passWd' and 'mnemonic'
fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
assertNotEqual(t, err, nil)
// get info of your account from OKChain
accInfo, err := okCli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
assertNotEqual(t, err, nil)
// transfer okb to addr1
res, err := okCli.Send(fromInfo, passWd, addr1, "10.24okt", "my memno", accInfo.GetAccountNumber(), accInfo.GetSequence())
assertNotEqual(t, err, nil)
```

you can use the object `okCli` to invoke more api functions.

### 6. Testing

All changes and addition of codes will be pushed with unit tests strictly. 

### 7. Contributing

No doubt that it's admirable to make contributions to OKChain Go SDK. You can provide your code as long as you have tested it with a local client and your unit test showed its validity.  

