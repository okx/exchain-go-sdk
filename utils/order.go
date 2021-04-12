package utils

import (
	"encoding/json"
	"errors"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ordertypes "github.com/okex/exchain/x/order/types"
)

// BuildOrderItems returns the set of OrderItem
// params must be checked by function CheckNewOrderParams
func BuildOrderItems(products, sides, prices, quantities []string) ([]ordertypes.OrderItem, error) {
	productsLen := len(products)
	orderItems := make([]ordertypes.OrderItem, productsLen)
	for i := 0; i < productsLen; i++ {
		priceDec, err := sdk.NewDecFromStr(prices[i])
		if err != nil {
			return nil, err
		}
		quantityDec, err := sdk.NewDecFromStr(quantities[i])
		if err != nil {
			return nil, err
		}

		orderItems[i] = ordertypes.OrderItem{
			Product:  products[i],
			Side:     sides[i],
			Price:    priceDec,
			Quantity: quantityDec,
		}
	}

	return orderItems, nil
}

// GetOrderIDsFromResponse filters the orderID from the tx response
// a useful tool
func GetOrderIDsFromResponse(txResp *sdk.TxResponse) (orderIDs []string, err error) {
	if len(txResp.Logs) != 1 {
		return orderIDs, errors.New("failed. only ONE msg could be in an order StdTx")
	}

	for _, event := range txResp.Logs[0].Events {
		if event.Type == "message" {
			for _, attribute := range event.Attributes {
				if attribute.Key == "orders" {
					var orderRes []ordertypes.OrderResult
					if err := json.Unmarshal([]byte(attribute.Value), &orderRes); err != nil {
						log.Println(ErrUnmarshalJSON(err.Error()).Error())
						continue
					}

					for _, res := range orderRes {
						orderIDs = append(orderIDs, res.OrderID)
					}
				}
			}
		}
	}

	return
}
