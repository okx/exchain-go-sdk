package types

// Module shows the expected behaviour of each module in okchain gosdk
type Module interface {
	RegisterCodec(cdc SDKCodec)
	Name() string
}
