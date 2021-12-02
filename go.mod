module github.com/okex/exchain-go-sdk

go 1.15

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/ethereum/go-ethereum v1.10.8
	github.com/golang/mock v1.5.0
	github.com/okex/exchain v0.19.17
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.2-0.20211120080923-5ccee577b032
	github.com/tendermint/tm-db => github.com/okex/tm-db v0.5.3-0.20211118074832-0d08619b97fe
)
