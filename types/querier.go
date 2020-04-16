package types




type MatchResult struct {
	Timestamp   int64   `gorm:"index;" json:"timestamp"`
	BlockHeight int64   `gorm:"PRIMARY_KEY;type:bigint" json:"block_height"`
	Product     string  `gorm:"PRIMARY_KEY;type:varchar(20)" json:"product"`
	Price       float64 `gorm:"type:DOUBLE" json:"price"`
	Quantity    float64 `gorm:"type:DOUBLE" json:"volume"`
}
