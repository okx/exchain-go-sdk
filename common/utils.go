package common

import (
	"errors"
	"fmt"
	"strings"
)

const (
	perPageDefault = 50
	perPageMax     = 200
)


func IsValidSide(side string) bool {
	if "BUY" != side && "SELL" != side {
		return false
	}
	return true
}

func CheckParamsPaging(start, end, page, perPage int) (perPageRet int, err error) {
	if start < 0 || end < 0 || page < 0 || perPage < 0 {
		return 0, errors.New("'start','end','page','perPage' cannot be negative")
	}

	if start > end {
		return 0, errors.New("'start' cannot be larger than 'end'")
	}

	if perPage == 0 {
		perPageRet = perPageDefault
	} else if perPage > perPageMax {
		perPageRet = perPageMax
	} else {
		perPageRet = perPage
	}
	return

}
