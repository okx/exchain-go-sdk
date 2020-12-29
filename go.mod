module github.com/okex/okexchain-go-sdk

go 1.15

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.9.25
	github.com/golang/mock v1.4.4
	github.com/okex/okexchain v0.11.1-0.20201229030024-5ebe8a154905
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.9
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-okexchain1.0.20201228082538-e886224efbda
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-okexchain1
)
