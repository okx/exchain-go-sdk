package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	cryptotypes "github.com/okex/exchain/libs/cosmos-sdk/crypto/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	ibcTypes "github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
)

// Ibc shows the expected behavior for inner ibc client
type Ibc interface {
	gosdktypes.Module
	IbcTx
	IbcQuery
}

// IbcTx send ibc tx
type IbcTx interface {

	// Transfer transfer token to destination chain
	Transfer(priKey cryptotypes.PrivKey, srcChannel string, receiver string, amount string, fee sdk.CoinAdapters, memo string, targetRpc string) (resp sdk.TxResponse, err error)
}

// IbcQuery shows the ibc query info
type IbcQuery interface {

	// QueryDenomTrace query a a denomination trace from a given hash.
	QueryDenomTrace(hash string) (*ibcTypes.QueryDenomTraceResponse, error)

	// QueryDenomTraces query all the denomination trace infos.
	QueryDenomTraces(page *query.PageRequest) (*ibcTypes.QueryDenomTracesResponse, error)

	// QueryIbcParams ibc-transfer parameter querying.
	QueryIbcParams() (*ibcTypes.QueryParamsResponse, error)

	// QueryEscrowAddress ibc-transfer parameter querying.
	QueryEscrowAddress(portID, channelID string) sdk.AccAddress
}
