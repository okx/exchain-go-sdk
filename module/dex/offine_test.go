package dex

import (
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/dex/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	unsignedPath = "./unsignedTx.json"
	signedPath   = "./signedTx.json"
)

func TestDexClient_GenerateUnsignedTransferOwnershipTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewDexClient(mockCli.MockBaseClient))

	fromAddr, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	toAddr, err := sdk.AccAddressFromBech32(recipient)
	require.NoError(t, err)

	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc)

	msg := types.NewMsgTransferOwnership(fromAddr, toAddr, product)
	expectedStdTx := sdk.NewStdTx([]sdk.Msg{msg}, sdk.NewStdFee(20000, mockCli.GetConfig().Fees), nil, memo)
	mockCli.EXPECT().BuildUnsignedStdTxOffline([]sdk.Msg{msg}, memo).Return(expectedStdTx)

	err = mockCli.Dex().GenerateUnsignedTransferOwnershipTx(product, addr, recipient, memo, unsignedPath)
	require.NoError(t, err)

	// read back to check
	stdTx, err := utils.GetStdTxFromFile(expectedCdc, unsignedPath)
	require.NoError(t, err)
	require.Equal(t, expectedStdTx, stdTx)

	// remove the temporary file
	err = os.Remove(unsignedPath)
	require.NoError(t, err)

}
