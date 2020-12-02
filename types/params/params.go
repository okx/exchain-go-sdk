package params

import (
	"time"
)

const (
	defaultPage    = 1
	defaultPerPage = 50
)

// QueryKlinesParams - structure of params to query the klines of a specific product
type QueryKlinesParams struct {
	Product     string
	Granularity int
	Size        int
}

// NewQueryKlinesParams creates a new instance of QueryKlinesParams
func NewQueryKlinesParams(product string, granularity, size int) QueryKlinesParams {
	return QueryKlinesParams{
		product,
		granularity,
		size,
	}
}

// QueryTickerParams - structure of params to query tickers
type QueryTickerParams struct {
	Product string `json:"product"`
	Count   int    `json:"count"`
	Sort    bool   `json:"sort"`
}

// NewQueryTickerParams creates a new instance of QueryTickerParams
func NewQueryTickerParams(product string, count int, sort bool) QueryTickerParams {
	return QueryTickerParams{
		product,
		count,
		sort,
	}
}

// QueryMatchParams - structure of params to query tx record
type QueryMatchParams struct {
	Product string
	Start   int64
	End     int64
	Page    int
	PerPage int
}

// NewQueryMatchParams creates a new instance of QueryMatchParams
func NewQueryMatchParams(product string, start, end int64, page, perPage int) QueryMatchParams {
	if page == 0 && perPage == 0 {
		page = defaultPage
		perPage = defaultPerPage
	}
	return QueryMatchParams{
		Product: product,
		Start:   start,
		End:     end,
		Page:    page,
		PerPage: perPage,
	}
}

// QueryOrderListParams - structure of params to query record list
type QueryOrderListParams struct {
	Address    string
	Product    string
	Page       int
	PerPage    int
	Start      int64
	End        int64
	Side       string
	HideNoFill bool
}

// NewQueryOrderListParams creates a new instance of NewQueryOrderListParams
func NewQueryOrderListParams(addr, product, side string, page, perPage int, start, end int64,
	hideNoFill bool) QueryOrderListParams {
	if page == 0 && perPage == 0 {
		page = defaultPage
		perPage = defaultPerPage
	}
	if start == 0 && end == 0 {
		end = time.Now().Unix()
	}
	return QueryOrderListParams{
		Address:    addr,
		Product:    product,
		Page:       page,
		PerPage:    perPage,
		Start:      start,
		End:        end,
		Side:       side,
		HideNoFill: hideNoFill,
	}
}

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
