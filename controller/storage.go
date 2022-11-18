package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const (
	KeyUid   = "user_id"
	KeyUname = "username"
	KeyToken = "token"
)

var UserNotLogin = errors.New("用户未登陆")

func GetCurrentUserId(ctx *gin.Context) (userId int64, err error) {
	uid, ok := ctx.Get(KeyUid)
	if !ok {
		err = UserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = UserNotLogin
		return
	}
	return
}

func GetCurrentUserToken(ctx *gin.Context) (token string, err error) {
	tk, ok := ctx.Get(KeyToken)
	if !ok {
		err = UserNotLogin
		return
	}
	token, ok = tk.(string)
	if !ok {
		err = UserNotLogin
		return
	}
	return
}
