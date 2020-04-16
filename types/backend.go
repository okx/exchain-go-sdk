package types






type Transaction struct {
	TxHash    string `gorm:"type:varchar(80)" json:"txhash"`
	Type      int64  `gorm:"index;" json:"type"` // 1:Transfer, 2:NewOrder, 3:CancelOrder
	Address   string `gorm:"index;type:varchar(80)" json:"address"`
	Symbol    string `gorm:"type:varchar(20)" json:"symbol"`
	Side      int64  `gorm:"" json:"side"` // 1:buy, 2:sell, 3:from, 4:to
	Quantity  string `gorm:"type:varchar(40)" json:"quantity"`
	Fee       string `gorm:"type:varchar(40)" json:"fee"`
	Timestamp int64  `gorm:"index" json:"timestamp"`
}