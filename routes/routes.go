package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"tiny-bbs/controller"
	"tiny-bbs/middleware"
)

func Init() *gin.Engine {
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"version": viper.GetString("app.version"),
		})
	})

	// 用户注册
	r.POST("/signup", controller.SignUpController)
	// 用户登陆
	r.POST("/login", controller.LoginController)
	// 主页
	r.GET("/home", middleware.JWTAuthMiddleware(), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"test_msg": "ok",
		})
	})
	return r
}
