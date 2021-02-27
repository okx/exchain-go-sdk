package evm

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/auth"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	"github.com/okex/okexchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvmClient_SendTx(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewEvmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().GetConfig().Return(gosdktypes.ClientConfig{Gas: 200000}).Times(3)
	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil).Times(3)

	res, err := mockCli.Evm().SendTx(fromInfo, passWd, recAddr, "0.1024", "", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	res, err = mockCli.Evm().SendTx(fromInfo, passWd, recAddrEth, "0.1024", "", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	res, err = mockCli.Evm().SendTx(fromInfo, passWd, recAddrEth, "0.1024", defaultPayloadStr[2:],
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Evm().SendTx(fromInfo, passWd, recAddr[1:], "0.1024", "", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Evm().SendTx(fromInfo, passWd, recAddr, "0.1024okt", "", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Evm().SendTx(fromInfo, passWd, recAddr, "0.1024", badPayloadStr, memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Evm().SendTx(fromInfo, passWd, recAddr, "0.1024okt", "", memo,
		accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestEvmClient_CreateContract(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewEvmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().GetConfig().Return(gosdktypes.ClientConfig{Gas: 200000}).Times(3)
	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil).Times(2)

	res, contractAddr, err := mockCli.Evm().CreateContract(fromInfo, passWd, "0.1024", defaultPayloadStr,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)
	require.NotZero(t, len(contractAddr))

	res, contractAddr, err = mockCli.Evm().CreateContract(fromInfo, passWd, "0.1024", defaultPayloadStr[2:],
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)
	require.NotZero(t, len(contractAddr))

	_, contractAddr, err = mockCli.Evm().CreateContract(fromInfo, passWd, "0.1024", badPayloadStr,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
	require.Zero(t, len(contractAddr))

	_, contractAddr, err = mockCli.Evm().CreateContract(fromInfo, passWd, "0.1024okt", defaultPayloadStr,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
	require.Zero(t, len(contractAddr))

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, contractAddr, err = mockCli.Evm().CreateContract(fromInfo, passWd, "0.1024", defaultPayloadStr,
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
	require.Zero(t, len(contractAddr))
}
