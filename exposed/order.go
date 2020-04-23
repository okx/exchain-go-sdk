package exposed

import (
	"github.com/okex/okchain-go-sdk/module/order/types"
	sdk "github.com/okex/okchain-go-sdk/types"
	"github.com/okex/okchain-go-sdk/types/crypto/keys"
)

// Order shows the expected behavior for inner order client
type Order interface {
	sdk.Module
	OrderTx
	OrderQuery
}

// OrderTx shows the expected tx behavior for inner order client
type OrderTx interface {
	NewOrders(fromInfo keys.Info, passWd, products, sides, prices, quantities, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
	CancelOrders(fromInfo keys.Info, passWd, orderIDs, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
}

// OrderQuery shows the expected query behavior for inner order client
type OrderQuery interface {
	QueryDepthBook(product string) (types.BookRes, error)
	QueryOrderDetail(orderID string) (types.OrderDetail, error)
}
