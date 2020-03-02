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
	ChangePercentage string `json:"changePercentage"`
}

type MatchResult struct {
	Timestamp   int64   `gorm:"index;" json:"timestamp"`
	BlockHeight int64   `gorm:"PRIMARY_KEY;type:bigint" json:"block_height"`
	Product     string  `gorm:"PRIMARY_KEY;type:varchar(20)" json:"product"`
	Price       float64 `gorm:"type:DOUBLE" json:"price"`
	Quantity    float64 `gorm:"type:DOUBLE" json:"volume"`
}
