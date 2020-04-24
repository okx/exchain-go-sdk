package backend

import (
	"errors"
	"testing"
	"time"

	"github.com/okex/okchain-go-sdk/module/backend/types"
	"github.com/okex/okchain-go-sdk/types/params"

	"github.com/golang/mock/gomock"
	"github.com/okex/okchain-go-sdk/mocks"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tendermint/libs/common"
)

const (
	addr    = "okchain1dcsxvxgj374dv3wt9szflf9nz6342juzzkjnlz"
	product = "btc-000_okt"
)

func TestBackendClient_QueryDeals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewBackendClient(mockCli.MockBaseClient))

	timestamp, blockHeight := time.Now().Unix()*1000, int64(1024)
	orderID, side, fee := "ID0000000000-1", "BUY", "0.00100000btc-000"
	price, quantity := 1024.1024, 2048.2048
	start, end, page, perPage := 0, 0, 1, 30

	expectedRet := mockCli.BuildBackendDealsResultBytes(timestamp, blockHeight, orderID, addr, product, side, fee, price, quantity)
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryDealsParams(addr, product, int64(start), int64(end), page, perPage, side)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(3)
	mockCli.EXPECT().Query(types.DealsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	deals, err := mockCli.Backend().QueryDeals(addr, product, side, start, end, page, perPage)
	require.NoError(t, err)
	require.Equal(t, timestamp, deals[0].Timestamp)
	require.Equal(t, blockHeight, deals[0].BlockHeight)
	require.Equal(t, orderID, deals[0].OrderID)
	require.Equal(t, addr, deals[0].Sender)
	require.Equal(t, product, deals[0].Product)
	require.Equal(t, side, deals[0].Side)
	require.Equal(t, price, deals[0].Price)
	require.Equal(t, quantity, deals[0].Quantity)
	require.Equal(t, fee, deals[0].Fee)

	_, err = mockCli.Backend().QueryDeals(addr[1:], product, side, start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, "", side, start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, product, "BUY&&SELL", start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, product, side, end+1, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, product, side, -1, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, product, side, start, -1, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, product, side, start, end, -1, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryDeals(addr, product, side, start, end, page, -1)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.DealsPath, cmn.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Backend().QueryDeals(addr, product, side, start, end, page, perPage)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.DealsPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Backend().QueryDeals(addr, product, side, start, end, page, perPage)
	require.Error(t, err)

}

func TestBackendClient_QueryOpenOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewBackendClient(mockCli.MockBaseClient))

	txHash, orderID, side := "default tx hash", "ID0000000000-1", "BUY"
	price, quantity, filledAvgQuantity, remainQuantity := "1024.1024", "2048.2048", "1024.1024", "4096.4096"
	status, timestamp := int64(0), time.Now().Unix()*1000
	start, end, page, perPage := 0, 0, 1, 30

	expectedRet := mockCli.BuildBackendOrdersResultBytes(txHash, orderID, addr, product, side, price, quantity, filledAvgQuantity,
		remainQuantity, status, timestamp)
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryOrderListParams(addr, product, side, page, perPage, int64(start), int64(end), false)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(3)
	mockCli.EXPECT().Query(types.OpenOrdersPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	openOrders, err := mockCli.Backend().QueryOpenOrders(addr, product, side, start, end, page, perPage)
	require.NoError(t, err)
	require.Equal(t, txHash, openOrders[0].TxHash)
	require.Equal(t, orderID, openOrders[0].OrderID)
	require.Equal(t, addr, openOrders[0].Sender)
	require.Equal(t, product, openOrders[0].Product)
	require.Equal(t, side, openOrders[0].Side)
	require.Equal(t, price, openOrders[0].Price)
	require.Equal(t, quantity, openOrders[0].Quantity)
	require.Equal(t, status, openOrders[0].Status)
	require.Equal(t, filledAvgQuantity, openOrders[0].FilledAvgPrice)
	require.Equal(t, remainQuantity, openOrders[0].RemainQuantity)
	require.Equal(t, timestamp, openOrders[0].Timestamp)

	_, err = mockCli.Backend().QueryOpenOrders(addr[1:], product, side, start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, "", side, start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, product, "BUY&&SELL", start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, end+1, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, -1, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, start, -1, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, start, end, -1, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, start, end, page, -1)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.OpenOrdersPath, cmn.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, start, end, page, perPage)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.OpenOrdersPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Backend().QueryOpenOrders(addr, product, side, start, end, page, perPage)
	require.Error(t, err)
}

func TestBackendClient_QueryClosedOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	config, err := sdk.NewClientConfig("testURL", "testChain", sdk.BroadcastBlock, "0.01okt", 200000)
	require.NoError(t, err)
	mockCli := mocks.NewMockClient(t, ctrl, config)
	mockCli.RegisterModule(NewBackendClient(mockCli.MockBaseClient))

	txHash, orderID, side := "default tx hash", "ID0000000000-1", "BUY"
	price, quantity, filledAvgQuantity, remainQuantity := "1024.1024", "2048.2048", "1024.1024", "4096.4096"
	status, timestamp := int64(0), time.Now().Unix()*1000
	start, end, page, perPage := 0, 0, 1, 30

	expectedRet := mockCli.BuildBackendOrdersResultBytes(txHash, orderID, addr, product, side, price, quantity, filledAvgQuantity,
		remainQuantity, status, timestamp)
	expectedCdc := mockCli.GetCodec()

	queryParams := params.NewQueryOrderListParams(addr, product, side, page, perPage, int64(start), int64(end), false)
	queryBytes := expectedCdc.MustMarshalJSON(queryParams)

	mockCli.EXPECT().GetCodec().Return(expectedCdc).Times(3)
	mockCli.EXPECT().Query(types.ClosedOrdersPath, cmn.HexBytes(queryBytes)).Return(expectedRet, nil)

	openOrders, err := mockCli.Backend().QueryClosedOrders(addr, product, side, start, end, page, perPage)
	require.NoError(t, err)
	require.Equal(t, txHash, openOrders[0].TxHash)
	require.Equal(t, orderID, openOrders[0].OrderID)
	require.Equal(t, addr, openOrders[0].Sender)
	require.Equal(t, product, openOrders[0].Product)
	require.Equal(t, side, openOrders[0].Side)
	require.Equal(t, price, openOrders[0].Price)
	require.Equal(t, quantity, openOrders[0].Quantity)
	require.Equal(t, status, openOrders[0].Status)
	require.Equal(t, filledAvgQuantity, openOrders[0].FilledAvgPrice)
	require.Equal(t, remainQuantity, openOrders[0].RemainQuantity)
	require.Equal(t, timestamp, openOrders[0].Timestamp)

	_, err = mockCli.Backend().QueryClosedOrders(addr[1:], product, side, start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, "", side, start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, product, "BUY&&SELL", start, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, end+1, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, -1, end, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, start, -1, page, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, start, end, -1, perPage)
	require.Error(t, err)

	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, start, end, page, -1)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.ClosedOrdersPath, cmn.HexBytes(queryBytes)).Return(expectedRet, errors.New("default error"))
	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, start, end, page, perPage)
	require.Error(t, err)

	mockCli.EXPECT().Query(types.ClosedOrdersPath, cmn.HexBytes(queryBytes)).Return(expectedRet[1:], nil)
	_, err = mockCli.Backend().QueryClosedOrders(addr, product, side, start, end, page, perPage)
	require.Error(t, err)
}
