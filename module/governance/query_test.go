package governance

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	govtypes "github.com/okex/okexchain/x/gov/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
	"time"
)

const (
	addr      = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okexchainpub17weu6qepq0ph2t3u697qar7rmdtdtqp4744jcprjd2h356zr0yh5vmw38a3my4vqjx5"
	mnemonic  = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo      = "my memo"
)

func TestGovClient_QueryProposals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewGovClient(mockCli.MockBaseClient))

	proposalID, status, mockTime := uint64(1024), govtypes.ProposalStatus(0x01), time.Now()
	mockPower, err := sdk.NewDecFromStr("0.25")
	require.NoError(t, err)
	totalDeposit, err := sdk.ParseDecCoins("1.024okt,2.048btc")
	require.NoError(t, err)

	var depositorAddr, voterAddr sdk.AccAddress
	var proposalStatus govtypes.ProposalStatus

	expectedRet := mockCli.BuildProposalsBytes(proposalID, status, mockTime, totalDeposit, mockPower)
	expectedCdc := mockCli.GetCodec()
	expectedParams := expectedCdc.MustMarshalJSON(govtypes.NewQueryProposalsParams(proposalStatus, 0, voterAddr, depositorAddr))
	expectedPath := fmt.Sprintf("custom/%s/proposals", govtypes.QuerierRoute)
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(9)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

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

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Governance().QueryProposals("", "", "", 0)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Governance().QueryProposals("", "", "", 0)
	require.Error(t, err)

	depositorAddr, err = sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	voterAddr, err = sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	_, err = mockCli.Governance().QueryProposals(addr[1:], addr, "deposit_period", 1)
	require.Error(t, err)

	_, err = mockCli.Governance().QueryProposals(addr, addr[1:], "deposit_period", 1)
	require.Error(t, err)

	expectedParams = expectedCdc.MustMarshalJSON(govtypes.NewQueryProposalsParams(govtypes.StatusDepositPeriod, 1, voterAddr, depositorAddr))
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	proposals, err = mockCli.Governance().QueryProposals(addr, addr, "deposit_period", 1)
	require.NoError(t, err)

	expectedParams = expectedCdc.MustMarshalJSON(govtypes.NewQueryProposalsParams(govtypes.StatusNil, 1, voterAddr, depositorAddr))
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	_, err = mockCli.Governance().QueryProposals(addr, addr, "unknown status", 1)
	require.NoError(t, err)
}
