package order

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	"github.com/okex/okchain-go-sdk/module/order/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"

	"testing"
)

const (
	addr      = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	name      = "alice"
	passWd    = "12345678"
	accPubkey = "okchainpub1addwnpepqgzuks5c07kfce85e0t0x8qkuvvxu874965ruafn6svhjrhswt0lgdj85lv"
	mnemonic  = "dumb thought reward exhibit quick manage force imitate blossom vendor ketchup sniff"
	memo      = "my memo"

	product = "btc-000_okt"
)

func TestOrderClient_QueryOrderDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt")
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

	mockCli.EXPECT().GetCodec().Return(expectedCdc)
	mockCli.EXPECT().Query(fmt.Sprintf("%s/%s", types.OrderDetailPath, orderID), nil).Return(expectedRet, nil)

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

}
