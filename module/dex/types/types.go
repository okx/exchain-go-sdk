package types

import (
	sdk "github.com/okex/okexchain-go-sdk/types"
)

// const
const (
	ModuleName = "dex"

	ProductsPath = "custom/dex/products"
)

var (
	msgCdc = sdk.NewCodec()
)

func init() {
	sdk.RegisterBasicCodec(msgCdc)
	RegisterCodec(msgCdc)
}

// RegisterCodec registers the msg type for dex module
func RegisterCodec(cdc sdk.SDKCodec) {
	cdc.RegisterConcrete(MsgList{}, "okchain/dex/MsgList")
	cdc.RegisterConcrete(MsgDeposit{}, "okchain/dex/MsgDeposit")
	cdc.RegisterConcrete(MsgWithdraw{}, "okchain/dex/MsgWithdraw")
	cdc.RegisterConcrete(MsgTransferOwnership{}, "okchain/dex/MsgTransferTradingPairOwnership")
	cdc.RegisterConcrete(MsgCreateOperator{}, "okchain/dex/CreateOperator")
	cdc.RegisterConcrete(MsgUpdateOperator{}, "okchain/dex/UpdateOperator")
}

// TokenPair represents token pair object
type TokenPair struct {
	BaseAssetSymbol  string         `json:"base_asset_symbol"`
	QuoteAssetSymbol string         `json:"quote_asset_symbol"`
	InitPrice        sdk.Dec        `json:"price"`
	MaxPriceDigit    int64          `json:"max_price_digit"`
	MaxQuantityDigit int64          `json:"max_size_digit"`
	MinQuantity      sdk.Dec        `json:"min_trade_size"`
	ID               uint64         `json:"token_pair_id"`
	Delisting        bool           `json:"delisting"`
	Owner            sdk.AccAddress `json:"owner"`
	Deposits         sdk.DecCoin    `json:"deposits"`
	BlockHeight      int64          `json:"block_height"`
}
