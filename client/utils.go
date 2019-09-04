package client

import "errors"

const (
	countDefault = 100
)

func checkParamsGetTickersInfo(count []int) (countRet int, err error) {
	if len(count) > 1 {
		return 0, errors.New("invalid params input for 'GetTickersInfo'")
	}

	if len(count) == 0 {
		countRet = countDefault
	} else {
		if count[0] < 0 {
			return 0, errors.New("'count' cannot be negative")
		}
		countRet = count[0]
	}
	return
}
