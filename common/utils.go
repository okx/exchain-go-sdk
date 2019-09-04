package common

import (
	"fmt"
	"github.com/ok-chain/gosdk/common/libs/pkg/errors"
	"strings"
)

func IsValidAccaddr(addr string) bool {
	if len(addr) != 46 || !strings.HasPrefix(addr, "okchain") {
		fmt.Println("address inputed is not valid")
		return false
	}
	return true
}

func CheckParamsGetTickersInfo(count int) error {
	if count < 0 {
		return errors.New("'count' cannot be negative")
	}
	return nil
}
