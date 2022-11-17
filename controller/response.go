package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
可以根据自己需要去定义和使用
*/

type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

//const (
//	CodeSuccess       = http.StatusOK
//	CodeInvalidParams = http.StatusBadRequest
//	CodeUserExist     = http.StatusGone
//	CodeUserNotExist  = http.StatusNotFound
//	CodeErrUserPsw    = http.StatusConflict
//	CodeServerBusy    = http.StatusInternalServerError
//)

const (
	CodeSuccess = 10000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeErrUserPsw
	CodeServerBusy
)

var CodeMsgText = map[int]string{
	CodeSuccess:       "success",
	CodeInvalidParams: "请求参数错误 ",
	CodeUserExist:     "用户已经存在",
	CodeUserNotExist:  "用户不存在",
	CodeErrUserPsw:    "密码或者用户名输入有误",
	CodeServerBusy:    "服务器繁忙",
}

func ResponseErr(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  CodeMsgText[code],
		Data: nil,
	})
}

func ResponseErrWithMsg(ctx *gin.Context, code int, msgErr string) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  CodeMsgText[code] + msgErr,
		Data: nil,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Msg:  CodeMsgText[CodeSuccess],
		Data: data,
	})
}
