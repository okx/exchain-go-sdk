package ibc

import (
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
)

func (ibc ibcClient) QueryDenomTrace(hash string) (types.QueryDenomTraceResponse, error) {

	return types.QueryDenomTraceResponse{}, nil
}

func (ibc ibcClient) QeuryDenomTraces(page *query.PageRequest) (types.QueryDenomTracesResponse, error) {
	return types.QueryDenomTracesResponse{}, nil
}
