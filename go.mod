module github.com/okex/exchain-go-sdk

go 1.16

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/ethereum/go-ethereum v1.10.8
	github.com/golang/mock v1.5.0
	github.com/okex/exchain v1.1.2
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.8-oec2
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-exchain2
	github.com/tendermint/tm-db => github.com/okex/tm-db v0.5.2-exchain4
)
