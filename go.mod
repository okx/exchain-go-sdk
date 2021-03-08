module github.com/okex/okexchain-go-sdk

go 1.15

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.9.25
	github.com/golang/mock v1.5.0
	github.com/okex/okexchain v0.16.8
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.33.9
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-okexchain7
	github.com/tendermint/iavl => github.com/okex/iavl v0.14.1-okexchain1
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-okexchain4
)
