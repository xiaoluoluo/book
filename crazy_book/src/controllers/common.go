package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type Response struct {
	Err     string `json:"err"`
	ErrCode uint32 `json:"errCode"`
	Data    string `json:"data"`
}

func buildResponse(err string, errCode uint32, data string) string {
	resp := &Response{
		Err:     err,
		ErrCode: errCode,
		Data:    data,
	}
	r, _ := json.Marshal(resp)
	return string(r)
}
func BuildErrResponse(err string) string {
	return buildResponse(err, 1, "")
}

func BuildSuccessResponse(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		logs.Error("Login.Marshal err:", err.Error())
		return BuildErrResponse("Marshal err")
	}
	return buildResponse("", 0, string(jsonData))
}
