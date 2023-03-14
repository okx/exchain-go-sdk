module github.com/okex/exchain-go-sdk

go 1.16

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/ethereum/go-ethereum v1.10.25
	github.com/golang/mock v1.6.0
	github.com/okx/okbchain v0.0.0-20230314082628-432e974ddf9e
	github.com/stretchr/testify v1.8.0
)

replace (
	github.com/ethereum/go-ethereum => github.com/okex/go-ethereum v1.10.26-0.20230313021040-8e34a70661c0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/go-amino => github.com/okex/go-amino v0.15.1-okc4
)
