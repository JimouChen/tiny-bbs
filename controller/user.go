package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tiny-bbs/dao/mysql"
	"tiny-bbs/models"
	"tiny-bbs/pkg/selfpkg"
	"tiny-bbs/service"
)

func SignUpController(ctx *gin.Context) {
	//- 参数获取和校验
	//var userMsg *models.ParmaRegister
	userMsg := new(models.ParmaRegister)
	if err := ctx.ShouldBindJSON(userMsg); err != nil {
		zap.L().Error("sign up with invalid param", zap.Error(err))
		ResponseErrWithMsg(ctx, CodeInvalidParams, selfpkg.Trans2cnForSignUp(err.Error()))
		return
	}
	// 手动判断亲请求参数合理性
	//if len(userMsg.RePassword) == 0 || len(userMsg.Password) == 0 || len(userMsg.Username) == 0 || userMsg.Password != userMsg.RePassword {
	//	zap.L().Error("sign up with invalid param")
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	//- 业务处理，调service层逻辑代码，service调dao层
	if err := service.SignUp(userMsg); err != nil {
		if errors.Is(err, mysql.ErrUserExist) {
			ResponseErrWithMsg(ctx, CodeUserExist, " 注册失败")
			return
		}
		ResponseErrWithMsg(ctx, CodeServerBusy, " 注册失败")

		return
	}
	//- 返回响应
	ResponseSuccess(ctx, "注册成功")
	return
}

func LoginController(ctx *gin.Context) {
	// 获取参数，去数据库比对
	userMsg := new(models.ParmaLogin)
	if err := ctx.ShouldBindJSON(userMsg); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		ResponseErrWithMsg(ctx, CodeInvalidParams, selfpkg.Trans2cnForSignUp(err.Error()))
		return
	}

	// 处理逻辑
	token, err := service.Login(userMsg)
	if err != nil {
		msgAdd := " 登陆失败"
		if errors.Is(err, mysql.ErrServerBusy) {
			ResponseErrWithMsg(ctx, CodeServerBusy, msgAdd)
		} else if errors.Is(err, mysql.ErrUserNotExist) {
			ResponseErrWithMsg(ctx, CodeUserNotExist, msgAdd)
		} else if errors.Is(err, mysql.ErrPswUName) {
			ResponseErrWithMsg(ctx, CodeErrUserPsw, msgAdd)
		}
		return
	}
	ctx.Set(KeyToken, token)
	ResponseSuccess(ctx, gin.H{"token": token})
}
