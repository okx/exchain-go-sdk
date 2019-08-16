package common

type BaseResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	DetailMsg string      `json:"detailMsg"`
	Data      interface{} `json:"data"`
}