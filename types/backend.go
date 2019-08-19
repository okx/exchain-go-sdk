package types

type Order struct {
	TxHash         string `gorm:"type:varchar(80)" json:"txHash"`
	OrderId        string `gorm:"PRIMARY_KEY;type:varchar(30)" json:"orderId"`
	Sender         string `gorm:"index;type:varchar(80)" json:"sender"`
	Product        string `gorm:"index;type:varchar(20)" json:"product"`
	Side           string `gorm:"type:varchar(10)" json:"side"`
	Price          string `gorm:"type:varchar(40)" json:"price"`
	Quantity       string `gorm:"type:varchar(40)" json:"quantity"`
	Status         int64  `gorm:"index;" json:"status"`
	FilledAvgPrice string `gorm:"type:varchar(40)" json:"filledAvgPrice"`
	RemainQuantity string `gorm:"type:varchar(40)" json:"remainQuantity"`
	Timestamp      int64  `gorm:"index;" json:"timestamp"`
}

type Deal struct {
	Timestamp   int64   `gorm:"index;" json:"timestamp"`
	BlockHeight int64   `gorm:"PRIMARY_KEY;type:bigint" json:"blockHeight"`
	OrderId     string  `gorm:"PRIMARY_KEY;type:varchar(30)" json:"orderId"`
	Sender      string  `gorm:"index;type:varchar(80)" json:"sender"`
	Product     string  `gorm:"index;type:varchar(20)" json:"product"`
	Side        string  `gorm:"type:varchar(10)" json:"side"`
	Price       float64 `gorm:"type:DOUBLE" json:"price"`
	Quantity    float64 `gorm:"type:DOUBLE" json:"volume"`
	Fee         string  `gorm:"type:varchar(20)" json:"fee"`
}


type Transaction struct {
	TxHash    string `gorm:"type:varchar(80)" json:"txHash"`
	Type      int64  `gorm:"index;" json:"type"` // 1:Transfer, 2:NewOrder, 3:CancelOrder
	Address   string `gorm:"index;type:varchar(80)" json:"address"`
	Symbol    string `gorm:"type:varchar(20)" json:"symbol"`
	Side      int64  `gorm:"" json:"side"` // 1:buy, 2:sell, 3:from, 4:to
	Quantity  string `gorm:"type:varchar(40)" json:"quantity"`
	Fee       string `gorm:"type:varchar(40)" json:"fee"`
	Timestamp int64  `gorm:"index" json:"timestamp"`
}