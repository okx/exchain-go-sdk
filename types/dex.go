package types

// TokenPair represents token pair object
type TokenPair struct {
	BaseAssetSymbol  string     `json:"base_asset_symbol"`
	QuoteAssetSymbol string     `json:"quote_asset_symbol"`
	InitPrice        Dec        `json:"price"`
	MaxPriceDigit    int64      `json:"max_price_digit"`
	MaxQuantityDigit int64      `json:"max_size_digit"`
	MinQuantity      Dec        `json:"min_trade_size"`
	ID               uint64     `json:"token_pair_id"`
	Delisting        bool       `json:"delisting"`
	Owner            AccAddress `json:"owner"`
	Deposits         DecCoin    `json:"deposits"`
	BlockHeight      int64      `json:"block_height"`
}
