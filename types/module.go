package types

// Module shows the expected behaviour of each module in okexchain gosdk
type Module interface {
	RegisterCodec(cdc SDKCodec)
	Name() string
}
