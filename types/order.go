package types

type OrderItem struct {
	Product  string `json:"product"`
	Side     string `json:"side"`
	Price    Dec    `json:"price"`
	Quantity Dec    `json:"quantity"`
}

// NewOrderItem creates a new instance of OrderItem
func NewOrderItem(product string, side string, price string, quantity string) OrderItem {
	return OrderItem{
		Product:  product,
		Side:     side,
		Price:    MustNewDecFromStr(price),
		Quantity: MustNewDecFromStr(quantity),
	}
}

// BuildOrderItems returns the set of OrderItem
// params must be checked by function CheckNewOrderParams
func BuildOrderItems(products, sides, prices, quantities []string) []OrderItem {
	productsLen := len(products)
	orderItems := make([]OrderItem, productsLen)
	for i := 0; i < productsLen; i++ {
		orderItems[i] = NewOrderItem(products[i], sides[i], prices[i], quantities[i])
	}

	return orderItems
}

type OrderResult struct {
	Code    uint32 `json:"code"`
	Message string `json:"msg"`
	OrderID string `json:"orderid"`
}
