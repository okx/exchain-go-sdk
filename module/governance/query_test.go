package governance

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/governance/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/params"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
	"testing"
	"time"
)

const (
	addr      = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
	mnemonic  = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	memo      = "my memo"
)

func TestGovClient_QueryProposals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient))

	proposalID, status, mockTime := uint64(1024), ProposalStatus(0x01), time.Now()
	mockPower, err := sdk.NewDecFromStr("0.25")
	require.NoError(t, err)
	totalDeposit, err := sdk.ParseDecCoins("1.024okt,2.048btc")
	require.NoError(t, err)

	expectedRet := mockCli.BuildProposalsBytes(proposalID, status, mockTime, totalDeposit, mockPower)
	expectedCdc := mockCli.GetCodec()

	var depositorAddr, voterAddr sdk.AccAddress
	var proposalStatus types.ProposalStatus
	queryParams := params.NewQueryProposalsParams(proposalStatus, 0, voterAddr, depositorAddr)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(types.ProposalsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	proposals, err := mockCli.Governance().QueryProposals("", "", "", 0)
	require.NoError(t, err)

	require.Equal(t, 1, len(proposals))
	require.Equal(t, "", proposals[0].GetTitle())
	require.Equal(t, "", proposals[0].GetDescription())
	require.Equal(t, proposalID, proposals[0].ProposalID)
	require.Equal(t, status, proposals[0].Status)
	require.Equal(t, totalDeposit, proposals[0].TotalDeposit)
	require.Equal(t, mockPower, proposals[0].FinalTallyResult.TotalPower)
	require.Equal(t, mockPower, proposals[0].FinalTallyResult.TotalVotedPower)
	require.Equal(t, mockPower, proposals[0].FinalTallyResult.Yes)
	require.Equal(t, mockPower, proposals[0].FinalTallyResult.Abstain)
	require.Equal(t, mockPower, proposals[0].FinalTallyResult.No)
	require.Equal(t, mockPower, proposals[0].FinalTallyResult.NoWithVeto)
	require.True(t, mockTime.Equal(proposals[0].SubmitTime))
	require.True(t, mockTime.Equal(proposals[0].DepositEndTime))
	require.True(t, mockTime.Equal(proposals[0].VotingStartTime))
	require.True(t, mockTime.Equal(proposals[0].VotingEndTime))

	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(types.ProposalsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Governance().QueryProposals("", "", "", 0)
	require.Error(t, err)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(types.ProposalsPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Governance().QueryProposals("", "", "", 0)
	require.Error(t, err)

	depositorAddr,err=sdk.AccAddressFromBech32(addr)
	require.NoError(t,err)
	voterAddr,err=sdk.AccAddressFromBech32(addr)
	require.NoError(t,err)
	queryParams = params.NewQueryProposalsParams(types.StatusDepositPeriod, 1, voterAddr, depositorAddr)
	queryBytes = expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(types.ProposalsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	proposals, err = mockCli.Governance().QueryProposals(addr, addr, "deposit_period", 1)
	require.NoError(t, err)

}
