module github.com/okex/okexchain-go-sdk

go 1.14

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/golang/mock v1.4.3
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/common v0.6.0 // indirect
	github.com/prometheus/procfs v0.0.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190704165056-9c2d0518ed81 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.9
	github.com/tendermint/tm-db v0.5.1

)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.3-0.20201126083234-87b90b625ac4
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.32.10-okexchain
)
