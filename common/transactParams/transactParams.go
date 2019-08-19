package transactParams

import (
	"fmt"
	"github.com/ok-chain/ok-gosdk/crypto/keys"
	"strings"
)

func IsValidSendParams(fromInfo keys.Info, passWd, toAddr string) bool {
	if fromInfo == nil {
		fmt.Println("input invalid keys info")
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

func IsValidNewOrderParams(fromInfo keys.Info, passWd, product, side, price, quantity, memo string) bool {
	if fromInfo == nil {
		fmt.Println("input invalid keys info")
		return false
	}
	if len(passWd) == 0 {
		fmt.Println("no password input")
		return false
	}
	if len(product) == 0 {
		fmt.Println("no product input")
		return false
	}
	if side != "BUY" && side != "SELL" {
		fmt.Println("side can only be \"BUY\" or \"SELL\"")
		return false
	}
	if !checkAccuracyOfStr(price, 1) {
		fmt.Println("input invalid price")
		return false
	}
	if !checkAccuracyOfStr(quantity, 2) {
		fmt.Println("input invalid quantity")
		return false
	}
	return true
}

func IsValidCancelOrderParams(fromInfo keys.Info, passWd string) bool {
	if fromInfo == nil {
		fmt.Println("input invalid keys info")
		return false
	}
	if len(passWd) == 0 {
		fmt.Println("no password input")
		return false
	}
	return true
}

func IsValidMultiSend(fromInfo keys.Info, passWd, transferStr string) bool {
	if fromInfo == nil {
		fmt.Println("input invalid keys info")
		return false
	}
	if len(passWd) == 0 {
		fmt.Println("no password input")
		return false
	}

	if len(transferStr) == 0 {
		fmt.Println("no transfer string input")
		return false
	}
	return true
}

func IsValidMint(fromInfo keys.Info, passWd, symbol string, amount int64) bool {
	if fromInfo == nil {
		fmt.Println("input invalid keys info")
		return false
	}
	if len(passWd) == 0 {
		fmt.Println("no password input")
		return false
	}
	if len(symbol) == 0 {
		fmt.Println("no symbol input")
		return false
	}
	if amount<0 {
		fmt.Println("input invalid amount. It should be positive.")
		return false
	}
	return true
}
func checkAccuracyOfStr(num string, accuracy int) bool {
	num = strings.TrimSpace(num)
	strs := strings.Split(num, ".")
	if len(strs) > 2 || len(strs) == 0 {
		return false
	} else if len(strs) == 2 {
		for i, v := range strs[1] {
			if i > accuracy-1 && v != '0' {
				fmt.Printf("the accuracy can't be larger than %d\n", accuracy)
				return false
			}
		}
	}
	return true
}
