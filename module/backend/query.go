package backend

import (
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/module/backend/types"
	"github.com/okex/okchain-go-sdk/utils"
)

func (bc backendClient) QueryCandles(product string, granularity, size int) (candles [][]string, err error) {
	klinesParams := params.NewQueryKlinesParams(product, granularity, size)
	jsonBytes, err := bc.GetCodec().MarshalJSON(klinesParams)
	if err != nil {
		return candles, utils.ErrMarshalJSON(err.Error())
	}

	res, err := bc.Query(types.CandlesPath, jsonBytes)
	if err != nil {
		return candles, utils.ErrClientQuery(err.Error())
	}

	if err = utils.GetDataFromBaseResponse(res, &candles); err != nil {
		return candles, utils.ErrFilterDataFromBaseResponse("candles", err.Error())
	}

	return
}
