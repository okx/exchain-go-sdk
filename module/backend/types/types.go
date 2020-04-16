package types

// const
const (
	ModuleName = "backend"

	CandlesPath = "custom/backend/candles"
	TickersPath = "custom/backend/tickers"
)

// Ticker - structure of ticker's detail data
type Ticker struct {
	Symbol           string `json:"symbol"`
	Product          string `json:"product"`
	Timestamp        string `json:"timestamp"`
	Open             string `json:"open"`
	Close            string `json:"close"`
	High             string `json:"high"`
	Low              string `json:"low"`
	Price            string `json:"price"`
	Volume           string `json:"volume"`
	Change           string `json:"change"`
	ChangePercentage string `json:"change_percentage"`
}
