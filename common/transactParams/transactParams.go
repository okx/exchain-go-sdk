package transactParams

import (
	"fmt"
	"github.com/ok-chain/ok-gosdk/crypto/keys"
	"strings"
)

func IsValidSendParams(fromInfo keys.Info, passWd, toAddr string) bool {
	if fromInfo == nil {
		fmt.Println("input invalid name")
		return false
	}
	if len(passWd) == 0 {
		fmt.Println("no password input")
		return false
	}
	if len(toAddr) != 46 || !strings.HasPrefix(toAddr, "okchain") {
		fmt.Println("input invalid receiver address")
		return false
	}
	return true
}
