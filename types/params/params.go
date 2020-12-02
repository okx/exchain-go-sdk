package params

const (
	defaultPage    = 1
	defaultPerPage = 50
)

// QueryDealsParams - structure of params to query the deals info of a specific product
type QueryDealsParams struct {
	Address string
	Product string
	Start   int64
	End     int64
	Page    int
	PerPage int
	Side    string
}

// NewQueryDealsParams creates a new instance of NewQueryDealsParams
func NewQueryDealsParams(addr, product string, start, end int64, page, perPage int, side string) QueryDealsParams {
	if page == 0 && perPage == 0 {
		page = defaultPage
		perPage = defaultPerPage
	}
	return QueryDealsParams{
		Address: addr,
		Product: product,
		Start:   start,
		End:     end,
		Page:    page,
		PerPage: perPage,
		Side:    side,
	}
}

// QueryTxListParams - structure of params to query the transaction info
type QueryTxListParams struct {
	Address   string
	TxType    int64
	StartTime int64
	EndTime   int64
	Page      int
	PerPage   int
}

// NewQueryTxListParams creates a new instance of QueryTxListParams
func NewQueryTxListParams(addr string, txType, startTime, endTime int64, page, perPage int) QueryTxListParams {
	if page == 0 && perPage == 0 {
		page = defaultPage
		perPage = defaultPerPage
	}
	return QueryTxListParams{
		Address:   addr,
		TxType:    txType,
		StartTime: startTime,
		EndTime:   endTime,
		Page:      page,
		PerPage:   perPage,
	}
}
