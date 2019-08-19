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
