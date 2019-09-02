package common

import (
	"fmt"
	"strings"
)

func IsValidAccaddr(addr string) bool {
	if len(addr) != 46 || !strings.HasPrefix(addr, "okchain"){
		fmt.Println("address inputed is not valid")
		return false
	}
	return true
}
