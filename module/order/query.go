package order

import (
	"github.com/okex/okchain-go-sdk/common/params"
	"github.com/okex/okchain-go-sdk/module/order/types"
	"github.com/okex/okchain-go-sdk/utils"
)

// QueryDepthBook gets the current depth book info of a specific product
func (oc orderClient) QueryDepthBook(product string) (depthBook types.BookRes, err error) {
	depthBookParams := params.NewQueryDepthBookParams(product, 200)
	jsonBytes, err := oc.GetCodec().MarshalJSON(depthBookParams)
	if err != nil {
		return depthBook, utils.ErrMarshalJSON(err.Error())
	}

	res, err := oc.Query(types.DepthbookPath, jsonBytes)
	if err != nil {
		return depthBook, utils.ErrClientQuery(err.Error())
	}

	if err = oc.GetCodec().UnmarshalJSON(res, &depthBook); err != nil {
		return depthBook, utils.ErrUnmarshalJSON(err.Error())
	}

	return
}
