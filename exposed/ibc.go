package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"
	cryptotypes "github.com/okex/exchain/libs/cosmos-sdk/crypto/types"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	ibcTypes "github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
)

// Auth shows the expected behavior for inner auth client
type Ibc interface {
	gosdktypes.Module
	IbcTx
	DenomTraceQuery
}

type IbcTx interface {
	Transfer(priKey cryptotypes.PrivKey, srcChannel string, receiver string, amount string, fee sdk.CoinAdapters, memo string, targetRpc string) (resp sdk.TxResponse, err error)
}

// DenomTraceQuery shows the denom trace info from a given trace hash
type DenomTraceQuery interface {
	QueryDenomTrace(hash string) (ibcTypes.QueryDenomTraceResponse, error)
	QeuryDenomTraces(page *query.PageRequest) (ibcTypes.QueryDenomTracesResponse, error)
}
