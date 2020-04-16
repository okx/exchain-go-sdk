package common


func IsValidSide(side string) bool {
	if "BUY" != side && "SELL" != side {
		return false
	}
	return true
}
