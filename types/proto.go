package types

type ProtocolDefinition struct {
	Version   uint64  `json:"version"`
	Software  string  `json:"software"`
	Height    uint64  `json:"height"`
	Threshold Dec `json:"threshold"`
}