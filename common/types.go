package common

type BaseResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detailMsg"`
	Data      interface{} `json:"data"`
}

type ParamPage struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
}

type ListDataRes struct {
	Data      interface{} `json:"data"`
	ParamPage ParamPage   `json:"paramPage"`
}

type ListResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detailMsg"`
	Data      ListDataRes `json:"data"`
}