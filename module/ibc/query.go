package ibc

import (
	"context"
	"crypto/sha256"
	"fmt"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
)

func (ibc ibcClient) QueryDenomTrace(hash string) (*types.QueryDenomTraceResponse, error) {
	req := &types.QueryDenomTraceRequest{
		Hash: hash,
	}

	return ibc.DenomTrace(context.Background(), req)
}

func (ibc ibcClient) QueryDenomTraces(page *query.PageRequest) (*types.QueryDenomTracesResponse, error) {
	req := &types.QueryDenomTracesRequest{
		Pagination: page,
	}

	return ibc.DenomTraces(context.Background(), req)
}

func (ibc ibcClient) QueryIbcParams() (*types.QueryParamsResponse, error) {
	req := &types.QueryParamsRequest{}

	return ibc.Params(context.Background(), req)
}

func (ibc ibcClient) QueryEscrowAddress(portID, channelID string) sdk.AccAddress {
	// a slash is used to create domain separation between port and channel identifiers to
	// prevent address collisions between escrow addresses created for different channels
	contents := fmt.Sprintf("%s/%s", portID, channelID)

	// ADR 028 AddressHash construction
	preImage := []byte(Version)
	preImage = append(preImage, 0)
	preImage = append(preImage, contents...)
	hash := sha256.Sum256(preImage)
	return hash[:20]
}
