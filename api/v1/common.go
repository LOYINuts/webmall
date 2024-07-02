package api

import (
	"encoding/json"
	"mywebmall/pkg/e"
	"mywebmall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: e.Error,
			Msg:    "json类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: e.Error,
		Msg:    "参数错误",
		Error:  err.Error(),
	}
}
