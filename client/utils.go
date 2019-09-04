package client

import "errors"

func checkParamsGetTickersInfo(count []int) (countRet int, err error) {
	if len(count) > 1 {
		return 0, errors.New("invalid params input for 'GetTickersInfo'")
	}

	if len(count) == 0 {
		countRet = 10
	} else {
		if count[0] < 0 {
			return 0, errors.New("'count' cannot be negative")
		}
		countRet = count[0]
	}
	return
}
