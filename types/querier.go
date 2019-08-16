package types

type BookRes struct {
	Asks []BookResItem `json:"asks"`
	Bids []BookResItem `json:"bids"`
}

type BookResItem struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type Tickers []Ticker

type Ticker struct {
	Symbol           string  `json:"symbol"`
	Product          string  `json:"product"`
	Timestamp        int64   `json:"timestamp"`
	Open             float64 `json:"open"`
	Close            float64 `json:"close"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Price            float64 `json:"price"`
	Volume           float64 `json:"volume"`
	Change           float64 `json:"change"`
	ChangePercentage string  `json:"changePercentage"`
}
