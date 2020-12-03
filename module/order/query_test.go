package order

import (
	"errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/okex/okexchain-go-sdk/mocks"
	gosdktypes "github.com/okex/okexchain-go-sdk/types"
	orderkeeper "github.com/okex/okexchain/x/order/keeper"
	ordertypes "github.com/okex/okexchain/x/order/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
)

const (
	addr      = "okexchain1ntvyep3suq5z7789g7d5dejwzameu08m6gh7yl"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okexchainpub17weu6qepq0ph2t3u697qar7rmdtdtqp4744jcprjd2h356zr0yh5vmw38a3my4vqjx5"
	mnemonic  = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	memo      = "my memo"

	product = "btc-000_okt"
)

func TestOrderClient_QueryOrderDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "", 200000,
		1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewOrderClient(mockCli.MockBaseClient))

	sender, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)
	price, err := sdk.NewDecFromStr("1.024")
	require.NoError(t, err)
	quantity, err := sdk.NewDecFromStr("1024.1024")
	require.NoError(t, err)
	filledAvgPrice, err := sdk.NewDecFromStr("1.024")
	require.NoError(t, err)
	remainQuantity, err := sdk.NewDecFromStr("10.24")
	require.NoError(t, err)
	remainLocked, err := sdk.NewDecFromStr("20.48")
	require.NoError(t, err)
	feePerBlock, err := sdk.ParseDecCoin("2.048okt")
	require.NoError(t, err)

	orderID := "ID0000000000-1"
	expectedRet := mockCli.BuildOrderDetailBytes("default txhash", orderID, "default extraInfo", product,
		"BUY", 0, 10240000, 1024, sender, price, quantity, filledAvgPrice, remainQuantity,
		remainLocked, feePerBlock)
	expectedCdc := mockCli.GetCodec()
	expectedPath := fmt.Sprintf("custom/%s/%s/%s", ordertypes.QuerierRoute, ordertypes.QueryOrderDetail, orderID)
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(2)
	mockCli.EXPECT().Query(expectedPath, nil).Return(expectedRet, int64(1024), nil)

	orderDetail, err := mockCli.Order().QueryOrderDetail(orderID)
	require.NoError(t, err)
	require.Equal(t, "default txhash", orderDetail.TxHash)
	require.Equal(t, orderID, orderDetail.OrderID)
	require.Equal(t, "default extraInfo", orderDetail.ExtraInfo)
	require.Equal(t, product, orderDetail.Product)
	require.Equal(t, feePerBlock, orderDetail.FeePerBlock)
	require.Equal(t, remainLocked, orderDetail.RemainLocked)
	require.Equal(t, remainQuantity, orderDetail.RemainQuantity)
	require.Equal(t, filledAvgPrice, orderDetail.FilledAvgPrice)
	require.Equal(t, quantity, orderDetail.Quantity)
	require.Equal(t, price, orderDetail.Price)
	require.Equal(t, sender, orderDetail.Sender)
	require.Equal(t, int64(1024), orderDetail.OrderExpireBlocks)
	require.Equal(t, int64(10240000), orderDetail.Timestamp)
	require.Equal(t, int64(0), orderDetail.Status)
	require.Equal(t, "BUY", orderDetail.Side)

	_, err = mockCli.Order().QueryOrderDetail("")
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, nil).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Order().QueryOrderDetail(orderID)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, nil).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Order().QueryOrderDetail(orderID)
	require.Error(t, err)
}

func TestOrderClient_QueryDepthBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := gosdktypes.NewClientConfig("testURL", "testChain", gosdktypes.BroadcastBlock, "",
		200000, 1.1, "0.00000001okt")
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewOrderClient(mockCli.MockBaseClient))

	expectedRet := mockCli.BuildBookResBytes("1.024", "10.24", "2.048", "20.48")
	expectedCdc := mockCli.GetCodec()
	expectedPath := fmt.Sprintf("custom/%s/%s", ordertypes.QuerierRoute, ordertypes.QueryDepthBook)
	expectedParams := expectedCdc.MustMarshalJSON(orderkeeper.NewQueryDepthBookParams(product, orderkeeper.DefaultBookSize))
	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(5)
	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet, int64(1024), nil)

	depthBook, err := mockCli.Order().QueryDepthBook(product)
	require.NoError(t, err)

	require.Equal(t, "1.024", depthBook.Asks[0].Price)
	require.Equal(t, "10.24", depthBook.Asks[0].Quantity)
	require.Equal(t, "2.048", depthBook.Bids[0].Price)
	require.Equal(t, "20.48", depthBook.Bids[0].Quantity)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(nil, int64(0), errors.New("default error"))
	_, err = mockCli.Order().QueryDepthBook(product)
	require.Error(t, err)

	mockCli.EXPECT().Query(expectedPath, tmbytes.HexBytes(expectedParams)).Return(expectedRet[1:], int64(1024), nil)
	_, err = mockCli.Order().QueryDepthBook(product)
	require.Error(t, err)
}
