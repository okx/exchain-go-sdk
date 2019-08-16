package queryParams

const (
	DefaultBookSize = 200
	DefaultPage    = 1
	DefaultPerPage = 50
)

type QueryAccTokenParams struct {
	Symbol string `json:"symbol"`
	Show   string `json:"show"`
}

func NewQueryAccTokenParams(symbol, show string) QueryAccTokenParams {
	return QueryAccTokenParams{
		symbol,
		show,
	}
}

type QueryDepthBookParams struct {
	Product string
	Size    int
}

func NewQueryDepthBookParams(product string, size int) QueryDepthBookParams {
	if size == 0 {
		size = DefaultBookSize
	}
	return QueryDepthBookParams{
		Product: product,
		Size:    size,
	}
}

type QueryKlinesParams struct {
	Product     string
	Granularity int
	Size        int
}

func NewQueryKlinesParams(product string, granularity, size int) QueryKlinesParams {
	return QueryKlinesParams{
		product,
		granularity,
		size,
	}
}

type QueryTickerParams struct {
	Product string `json:"product"`
	Count   int    `json:"count"`
	Sort    bool   `json:"sort"`
}

func NewQueryTickerParams(product string, count int, sort bool) QueryTickerParams {
	return QueryTickerParams{
		product,
		count,
		sort,
	}
}

type QueryMatchParams struct {
	Product string
	Start   int64
	End     int64
	Page    int
	PerPage int
}

func NewQueryMatchParams(product string, start, end int64, page, perPage int) QueryMatchParams {
	if page == 0 && perPage == 0 {
		page = DefaultPage
		perPage = DefaultPerPage
	}
	return QueryMatchParams{
		Product: product,
		Start:   start,
		End:     end,
		Page:    page,
		PerPage: perPage,
	}
}