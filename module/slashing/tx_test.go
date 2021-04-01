package slashing

import (
	"testing"

	gosdktypes "github.com/okex/okexchain-go-sdk/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	"github.com/okex/okexchain-go-sdk/module/auth"
	"github.com/okex/okexchain-go-sdk/utils"
	"github.com/stretchr/testify/require"
)

const (
	addr      = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "expub17weu6qepqtfc6zq8dukwc3lhlhx7th2csfjw0g3cqnqvanh7z9c2nhkr8mn5z9uq4q6"
	mnemonic  = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo      = "my memo"
)

func TestSlashingClient_Unjail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testchain-1", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewSlashingClient(mockCli.MockBaseClient), auth.NewAuthClient(mockCli.MockBaseClient))

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

	res, err := mockCli.Slashing().Unjail(fromInfo, passWd, memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.NoError(t, err)
	require.Equal(t, uint32(0), res.Code)

	_, err = mockCli.Slashing().Unjail(fromInfo, "", memo, accInfo.GetAccountNumber(), accInfo.GetSequence())
	require.Error(t, err)
}
