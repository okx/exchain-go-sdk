package common

type BaseResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	detail_msg string      `json:"detail_msg"`
	Data      interface{} `json:"data"`
}

type ParamPage struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

type ListDataRes struct {
	Data      interface{} `json:"data"`
	ParamPage ParamPage   `json:"param_page"`
}

type ListResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	detail_msg string      `json:"detail_msg"`
	Data      ListDataRes `json:"data"`
}