package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
		ctx.JSON(http.StatusOK, gin.H{
			"msg": selfpkg.Trans2cnForSignUp(err.Error()),
		})
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
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册失败:" + err.Error(),
		})
		return
	}
	//- 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功!",
	})
	return
}

func LoginController(ctx *gin.Context) {
	// 获取参数，去数据库比对
	userMsg := new(models.ParmaLogin)
	if err := ctx.ShouldBindJSON(userMsg); err != nil {
		fmt.Println(userMsg)

		zap.L().Error("login with invalid param", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": selfpkg.Trans2cnForSignUp(err.Error()),
		})
		return
	}

	// 处理逻辑
	if err := service.Login(userMsg); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "登陆失败： " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登陆成功",
	})
}
