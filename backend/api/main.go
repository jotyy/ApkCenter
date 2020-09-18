package api

import (
	"encoding/json"
	"file-upload/conf"
	"file-upload/serializer"
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// @Summary Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// @Summary ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Code: 40001,
				Msg:  fmt.Sprintf("%s%s", field, tag),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Code: 40001,
			Msg:  "JSON类型不匹配",
		}
	}

	return serializer.Response{
		Code: 40001,
		Msg:  "参数错误",
	}
}
