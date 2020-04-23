package utils

import (
	"encoding/json"
	"fmt"
	"github.com/okex/okchain-go-sdk/module/order/types"
	sdk "github.com/okex/okchain-go-sdk/types"
)

// GetOrderIDsFromResponse filters the orderID from the tx response
// a useful tool
func GetOrderIDsFromResponse(txResp *sdk.TxResponse) (orderIDs []string) {
	for _, event := range txResp.Events {
		if event.Type == "message" {
			for _, attribute := range event.Attributes {
				if attribute.Key == "orders" {
					var orderRes []types.OrderResult
					if err := json.Unmarshal([]byte(attribute.Value), &orderRes); err != nil {
						fmt.Println(err)
						return
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
