## OKChain Go SDK

The OKChain Go SDK seems like a useful lightweight Go library to interact with OKChain.The api functions of querying and transacting are exposed to the ones who installed the SDK. The Go SDK supports synchronous , asynchronous and blockmode requests.

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

The developer can install the OKChain Go SDK by `git clone` from github : xxxxxxxx(address in github)

### 3. API

The api functions of querying and transacting are in the files 'okclient/query,go'  and `okclient/transact.go`. You can find more details on the wiki : xxxxxxxxxx(okchain wiki address)

### 4. Node query

 The OKChain Go SDK also provides node query functions so that the underlying information of the blockchain can be got by developers.

The node query functions could be found in file `okclient/node_query.go `. The module of them are in file `okclient/node_query_test.go`. The develop could make it through and get clear that how to invoke them.

More details of node querying from OKChain on wiki : xxxxxxxxxx(okchain wiki, Part :node rpc address)

### 5. Example

Every operation by using Go SDK needs `okclient`. Here are the examples :

```go
// rpcUrl can be modified according to the actual situation
rpcUrl	 := "tcp://127.0.0.1:26657"
okCli 	 := NewClient(rpcUrl)
name     := "alice"
passWd   := "12345678"
mnemonic := "sustain hole urban away boy core lazy brick wait drive tiger tell"
addr1    := "okchain1dycww54mz20sfakx7hqtkf2ghdlx6tjry977gy"

// create your account key info by 'name','passWd' and 'mnemonic'
fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
assertNotEqual(t, err, nil)
// get info of your account from OKChain
accInfo, err := okCli.GetAccountInfoByAddr(fromInfo.GetAddress().String())
assertNotEqual(t, err, nil)
// transfer okb to addr1
res, err := okCli.Send(fromInfo, passWd, addr1, "10.24okb", "I love OK", accInfo.GetAccountNumber(), accInfo.GetSequence())
assertNotEqual(t, err, nil)
```

you can use the object `okCli` to invoke more api functions.

### 6. Testing

All changes and addition of codes will be pushed with unit tests strictly. 

### 7. Contributing

No doubt that it's admirable to make contributions to OKChain Go SDK. You can provide your code as long as you have tested it with a local client and your unit test showed its validity.  

