package controllers

import "encoding/json"

type Response struct {
	Err     string `json:"err"`
	ErrCode uint32 `json:"errCode"`
	Data    string `json:"data"`
}

func BuildResponse(err string, errCode uint32, data string) string {
	resp := &Response{
		Err:     err,
		ErrCode: errCode,
		Data:    data,
	}
	r, _ := json.Marshal(resp)
	return string(r)
}
func BuildErrResponse(err string) string {
	return BuildResponse(err, 1, "")
}
func BuildSuccessResponse(data string) string {
	return BuildResponse("", 0, data)
}
