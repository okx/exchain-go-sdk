module github.com/okex/exchain-go-sdk

go 1.15

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.10.8
	github.com/golang/mock v1.5.0
	github.com/okex/exchain v0.19.9
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.33.9
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-exchain17
	github.com/tendermint/iavl => github.com/okex/iavl v0.14.3-exchain2
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-exchain13
	github.com/tendermint/tm-db => github.com/okex/tm-db v0.5.2-exchain1
)
