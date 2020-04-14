package types

// useful for subspace and store query
var (
	ValidatorsKey = []byte{0x21}
	DelegatorKey        = []byte{0x52}
)

// GetDelegatorKey gets the key for Delegator
func GetDelegatorKey(delAddr AccAddress) []byte {
	return append(DelegatorKey, delAddr.Bytes()...)
}