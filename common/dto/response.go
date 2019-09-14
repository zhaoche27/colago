package dto

type Response struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	Data    interface{} `json:"data"`
}

func NewResponseOfSuccessWithoutData() *Response {
	return NewResponseOfSuccess(nil)
}

func NewResponseOfSuccess(data interface{}) *Response {
	return &Response{Data: data}
}

func NewResponseOfFailure(errCode int, errMsg string) *Response {
	return &Response{ErrCode: errCode, ErrMsg: errMsg}
}
