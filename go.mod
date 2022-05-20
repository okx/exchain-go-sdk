module github.com/okex/exchain-go-sdk

go 1.16

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/ethereum/go-ethereum v1.10.8
	github.com/golang/mock v1.6.0
	github.com/google/pprof v0.0.0-20220318212150-b2ab0324ddda // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20220319035150-800ac71e25c2 // indirect
	github.com/kisielk/godepgraph v0.0.0-20190626013829-57a7e4a651a9 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/minio/highwayhash v1.0.1 // indirect
	github.com/okex/exchain v1.3.2-0.20220519080729-7d16e4ca6148 // indirect
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/tendermint v0.34.14
	golang.org/x/net v0.0.0-20210903162142-ad29c8ab022f // indirect
)

replace (
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.8-oec3
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-exchain6
)
