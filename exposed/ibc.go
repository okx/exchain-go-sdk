package exposed

import (
	gosdktypes "github.com/okex/exchain-go-sdk/types"

	sdk "github.com/okx/okbchain/libs/cosmos-sdk/types"
	"github.com/okx/okbchain/libs/cosmos-sdk/types/query"
	ibcTypes "github.com/okx/okbchain/libs/ibc-go/modules/apps/transfer/types"
	client_types "github.com/okx/okbchain/libs/ibc-go/modules/core/02-client/types"
	chantypes "github.com/okx/okbchain/libs/ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/okx/okbchain/libs/ibc-go/modules/core/exported"
	tmcrypto "github.com/okx/okbchain/libs/tendermint/crypto"
	ctypes "github.com/okx/okbchain/libs/tendermint/rpc/core/types"
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
	Transfer(priKey tmcrypto.PrivKey, srcChannel string, receiver string, amount string, fee sdk.CoinAdapters, memo string, timeoutHeight client_types.Height) (resp sdk.TxResponse, err error)
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

	// QueryChannels query channels
	QueryChannels() (*chantypes.QueryChannelsResponse, error)

	// QueryTx query a tx
	QueryTx(hash string) (*ctypes.ResultTx, error)

	// QueryTxs
	QueryTxs(page, limit int, events []string) ([]*ctypes.ResultTx, error)

	// QueryHeaderAtHeight
	QueryHeaderAtHeight(height int64) (ibcexported.Header, error)

	// QueryChannel
	QueryChannel(req *chantypes.QueryChannelRequest) (*chantypes.QueryChannelResponse, error)

	// ConnectionChannels queries all the channels associated with a connection
	// end.
	ConnectionChannels(req *chantypes.QueryConnectionChannelsRequest) (*chantypes.QueryConnectionChannelsResponse, error)
	// ChannelClientState queries for the client state for the channel associated
	// with the provided channel identifiers.
	ChannelClientState(req *chantypes.QueryChannelClientStateRequest) (*chantypes.QueryChannelClientStateResponse, error)
	// ChannelConsensusState queries for the consensus state for the channel
	// associated with the provided channel identifiers.
	ChannelConsensusState(req *chantypes.QueryChannelConsensusStateRequest) (*chantypes.QueryChannelConsensusStateResponse, error)
	// PacketCommitment queries a stored packet commitment hash.
	PacketCommitment(req *chantypes.QueryPacketCommitmentRequest) (*chantypes.QueryPacketCommitmentResponse, error)
	// PacketCommitments returns all the packet commitments hashes associated
	// with a channel.
	PacketCommitments(req *chantypes.QueryPacketCommitmentsRequest) (*chantypes.QueryPacketCommitmentsResponse, error)
	// PacketReceipt queries if a given packet sequence has been received on the
	// queried chain
	PacketReceipt(req *chantypes.QueryPacketReceiptRequest) (*chantypes.QueryPacketReceiptResponse, error)
	// PacketAcknowledgement queries a stored packet acknowledgement hash.
	PacketAcknowledgement(req *chantypes.QueryPacketAcknowledgementRequest) (*chantypes.QueryPacketAcknowledgementResponse, error)
	// PacketAcknowledgements returns all the packet acknowledgements associated
	// with a channel.
	PacketAcknowledgements(req *chantypes.QueryPacketAcknowledgementsRequest) (*chantypes.QueryPacketAcknowledgementsResponse, error)
	// UnreceivedPackets returns all the unreceived IBC packets associated with a
	// channel and sequences.
	UnreceivedPackets(req *chantypes.QueryUnreceivedPacketsRequest) (*chantypes.QueryUnreceivedPacketsResponse, error)
	// UnreceivedAcks returns all the unreceived IBC acknowledgements associated
	// with a channel and sequences.
	UnreceivedAcks(req *chantypes.QueryUnreceivedAcksRequest) (*chantypes.QueryUnreceivedAcksResponse, error)
	// NextSequenceReceive returns the next receive sequence for a given channel.
	NextSequenceReceive(req *chantypes.QueryNextSequenceReceiveRequest) (*chantypes.QueryNextSequenceReceiveResponse, error)
}
