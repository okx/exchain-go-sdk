package utils

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ordertypes "github.com/okex/exchain/x/order/types"
	"github.com/stretchr/testify/require"
)

const (
	defaultProduct = "usdk_okt"
	defaultSide    = "BUY"
	defaultDecStr  = "1024.1024"
)

func TestBuildOrderItems(t *testing.T) {
	expectedProductsStrs := []string{defaultProduct, defaultProduct}
	expectedSidesStrs := []string{defaultSide, defaultSide}
	expectedPricesStrs := []string{defaultDecStr, defaultDecStr}
	expectedQuantitiesStrs := []string{defaultDecStr, defaultDecStr}
	expectedDec := sdk.MustNewDecFromStr(defaultDecStr)

	orderItems, err := BuildOrderItems(expectedProductsStrs, expectedSidesStrs, expectedPricesStrs, expectedQuantitiesStrs)
	require.NoError(t, err)
	require.Equal(t, 2, len(orderItems))
	for _, item := range orderItems {
		require.Equal(t, defaultProduct, item.Product)
		require.Equal(t, defaultSide, item.Side)
		require.Equal(t, expectedDec, item.Price)
		require.Equal(t, expectedDec, item.Price)
	}

	badDecStr := "1.024.1024.1024"
	expectedPricesStrs = []string{badDecStr, defaultDecStr}
	_, err = BuildOrderItems(expectedProductsStrs, expectedSidesStrs, expectedPricesStrs, expectedQuantitiesStrs)
	require.Error(t, err)

	expectedPricesStrs = []string{defaultDecStr, defaultDecStr}
	expectedQuantitiesStrs = []string{badDecStr, defaultDecStr}
	_, err = BuildOrderItems(expectedProductsStrs, expectedSidesStrs, expectedPricesStrs, expectedQuantitiesStrs)
	require.Error(t, err)
}

func TestGetOrderIDsFromResponse(t *testing.T) {
	mockOrderIDs := []string{"ID0000000000-1", "ID0000000000-2", "ID0000000000-3", "ID0000000000-4", "ID0000000000-5"}
	var orderResults, fakeOrderResults1, fakeOrderResults2 []ordertypes.OrderResult
	for i, orderID := range mockOrderIDs {
		mockOrderRes := buildMockOrderRes(orderID)
		if i < 3 {
			orderResults = append(orderResults, mockOrderRes)
		} else if i == 3 {
			fakeOrderResults1 = append(fakeOrderResults1, mockOrderRes)
		} else {
			fakeOrderResults2 = append(fakeOrderResults2, mockOrderRes)
		}
	}

	rawStrs := getRawStrSlice(orderResults, fakeOrderResults1, fakeOrderResults2)
	rawStrs = append(rawStrs, "string that failed to unmarshal JSON")
	require.Equal(t, 4, len(rawStrs))

	// build mock TxResponse
	mockTxResp := sdk.TxResponse{
		Logs: sdk.ABCIMessageLogs{
			{
				MsgIndex: 0,
				Log:      "default log",
				Events: sdk.StringEvents{
					{
						Type: "message",
						Attributes: []sdk.Attribute{
							{
								Key:   "not orders",
								Value: rawStrs[1],
							},
							{
								Key:   "orders",
								Value: rawStrs[3], // log error
							},
							{
								Key:   "orders",
								Value: rawStrs[0], // target
							},
						},
					},
					{
						Type: "not message",
						Attributes: []sdk.Attribute{
							{
								Key:   "not orders",
								Value: rawStrs[1],
							},
							{
								Key:   "orders",
								Value: rawStrs[2],
							},
						},
					},
				},
			},
		},
	}

	orderIDs, err := GetOrderIDsFromResponse(&mockTxResp)
	require.NoError(t, err)
	require.Equal(t, mockOrderIDs[:3], orderIDs)
}

func getRawStrSlice(orderResults ...[]ordertypes.OrderResult) (strs []string) {
	for _, res := range orderResults {
		bytes, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		strs = append(strs, string(bytes))
	}

	return
}

func buildMockOrderRes(orderID string) ordertypes.OrderResult {
	return ordertypes.OrderResult{
		Message: "default message",
		OrderID: orderID,
	}
}
