package ibc

import (
	"context"
	"crypto/sha256"
	"fmt"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/types/query"
	"github.com/okex/exchain/libs/ibc-go/modules/apps/transfer/types"
	chantypes "github.com/okex/exchain/libs/ibc-go/modules/core/04-channel/types"
)

func (ibc ibcClient) QueryDenomTrace(hash string) (*types.QueryDenomTraceResponse, error) {
	req := &types.QueryDenomTraceRequest{
		Hash: hash,
	}
	out := new(types.QueryDenomTraceResponse)
	err := ibc.Invoke(context.Background(), "/ibc.applications.transfer.v1.Query/DenomTrace", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) QueryDenomTraces(page *query.PageRequest) (*types.QueryDenomTracesResponse, error) {
	req := &types.QueryDenomTracesRequest{
		Pagination: page,
	}

	out := new(types.QueryDenomTracesResponse)
	err := ibc.Invoke(context.Background(), "/ibc.applications.transfer.v1.Query/DenomTraces", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) QueryIbcParams() (*types.QueryParamsResponse, error) {
	req := &types.QueryParamsRequest{}

	out := new(types.QueryParamsResponse)
	err := ibc.Invoke(context.Background(), "/ibc.applications.transfer.v1.Query/Params", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (ibc ibcClient) QueryChannels() (*chantypes.QueryChannelsResponse, error) {
	req := &chantypes.QueryChannelsRequest{}
	out := new(chantypes.QueryChannelsResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/Channels", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) QueryChannel(req *chantypes.QueryChannelRequest) (*chantypes.QueryChannelResponse, error) {
	out := new(chantypes.QueryChannelResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/Channel", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) ConnectionChannels(req *chantypes.QueryConnectionChannelsRequest) (*chantypes.QueryConnectionChannelsResponse, error) {
	out := new(chantypes.QueryConnectionChannelsResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/ConnectionChannels", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) ChannelClientState(req *chantypes.QueryChannelClientStateRequest) (*chantypes.QueryChannelClientStateResponse, error) {
	out := new(chantypes.QueryChannelClientStateResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/ChannelClientState", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) ChannelConsensusState(req *chantypes.QueryChannelConsensusStateRequest) (*chantypes.QueryChannelConsensusStateResponse, error) {
	out := new(chantypes.QueryChannelConsensusStateResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/ChannelConsensusState", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) PacketCommitment(req *chantypes.QueryPacketCommitmentRequest) (*chantypes.QueryPacketCommitmentResponse, error) {
	out := new(chantypes.QueryPacketCommitmentResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/PacketCommitment", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) PacketCommitments(req *chantypes.QueryPacketCommitmentsRequest) (*chantypes.QueryPacketCommitmentsResponse, error) {
	out := new(chantypes.QueryPacketCommitmentsResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/PacketCommitments", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) PacketReceipt(req *chantypes.QueryPacketReceiptRequest) (*chantypes.QueryPacketReceiptResponse, error) {
	out := new(chantypes.QueryPacketReceiptResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/PacketReceipt", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) PacketAcknowledgement(req *chantypes.QueryPacketAcknowledgementRequest) (*chantypes.QueryPacketAcknowledgementResponse, error) {
	out := new(chantypes.QueryPacketAcknowledgementResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/PacketAcknowledgement", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) PacketAcknowledgements(req *chantypes.QueryPacketAcknowledgementsRequest) (*chantypes.QueryPacketAcknowledgementsResponse, error) {
	out := new(chantypes.QueryPacketAcknowledgementsResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/PacketAcknowledgements", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) UnreceivedPackets(req *chantypes.QueryUnreceivedPacketsRequest) (*chantypes.QueryUnreceivedPacketsResponse, error) {
	out := new(chantypes.QueryUnreceivedPacketsResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/UnreceivedPackets", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) UnreceivedAcks(req *chantypes.QueryUnreceivedAcksRequest) (*chantypes.QueryUnreceivedAcksResponse, error) {
	out := new(chantypes.QueryUnreceivedAcksResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/UnreceivedAcks", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (ibc ibcClient) NextSequenceReceive(req *chantypes.QueryNextSequenceReceiveRequest) (*chantypes.QueryNextSequenceReceiveResponse, error) {
	out := new(chantypes.QueryNextSequenceReceiveResponse)
	err := ibc.Invoke(context.Background(), "/ibc.core.channel.v1.Query/NextSequenceReceive", req, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
