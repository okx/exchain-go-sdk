package farm

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

const (
	defaultPoolName = "default pool name"
)

func TestFarmClient_CreatePool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Farm().CreatePool(fromInfo, passWd, defaultPoolName, "0.1024okt", "test-token2",
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Farm().CreatePool(fromInfo, passWd, "", "0.1024okt", "test-token2",
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().CreatePool(fromInfo, passWd, defaultPoolName, "", "test-token2",
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().CreatePool(fromInfo, passWd, defaultPoolName, "0.1024okt", "",
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().CreatePool(fromInfo, "", defaultPoolName, "0.1024okt", "test-token2",
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Farm().CreatePool(fromInfo, passWd, defaultPoolName, "0.1024okt", "test-token2",
		memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestFarmClient_DestroyPool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Farm().DestroyPool(fromInfo, passWd, defaultPoolName, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Farm().DestroyPool(fromInfo, passWd, "", memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().DestroyPool(fromInfo, "", defaultPoolName, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Farm().DestroyPool(fromInfo, passWd, defaultPoolName, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestFarmClient_Provide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Farm().Provide(fromInfo, passWd, defaultPoolName, "1024.1024okt", "2048.2048",
		1024, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Farm().Provide(fromInfo, passWd, "", "1024.1024okt", "2048.2048",
		1024, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Provide(fromInfo, "", defaultPoolName, "1024.1024okt", "2048.2048",
		1024, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Provide(fromInfo, passWd, defaultPoolName, "1024.1024", "2048.2048",
		1024, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Provide(fromInfo, passWd, defaultPoolName, "1024.1024okt", "",
		1024, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Farm().Provide(fromInfo, passWd, defaultPoolName, "1024.1024okt", "2048.2048",
		1024, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}

func TestFarmClient_Lock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Farm().Lock(fromInfo, passWd, defaultPoolName, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Farm().Lock(fromInfo, passWd, "", "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Lock(fromInfo, "", defaultPoolName, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Lock(fromInfo, passWd, defaultPoolName, "1024.1024", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Farm().Lock(fromInfo, passWd, defaultPoolName, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)
}

func TestFarmClient_Unlock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewFarmClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

	fromInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	require.NoError(t, err)

	accBytes := mockCli.BuildAccountBytes(addr, accPubkey, "", "1024okt", 1, 2)
	expectedCdc := mockCli.GetCodec()
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(gomock.Any(), gomock.Any()).Return(accBytes, int64(1024), nil)

	accInfo, err := mockCli.Auth().QueryAccount(addr)
	require.NoError(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(mocks.DefaultMockSuccessTxResponse(), nil)

	res, err := mockCli.Farm().Unlock(fromInfo, passWd, defaultPoolName, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Farm().Unlock(fromInfo, passWd, "", "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Unlock(fromInfo, "", defaultPoolName, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	_, err = mockCli.Farm().Unlock(fromInfo, passWd, defaultPoolName, "1024.1024", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)

	mockCli.EXPECT().BuildAndBroadcast(
		fromInfo.GetName(), passWd, memo, gomock.AssignableToTypeOf([]sdk.Msg{}), accInfo.GetAccountNumber(), accInfo.GetSequence()).
		Return(sdk.TxResponse{}, errors.New("default error"))
	_, err = mockCli.Farm().Unlock(fromInfo, passWd, defaultPoolName, "1024.1024okt", memo, accInfo.GetAccountNumber(),
		accInfo.GetSequence())
	require.Error(t, err)
}
