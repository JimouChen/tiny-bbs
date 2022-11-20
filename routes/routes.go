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

	v1 := r.Group("/api/v1")

	// 用户注册
	v1.POST("/signup", controller.SignUpController)
	// 用户登陆
	v1.POST("/login", controller.LoginController)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityController)
		v1.GET("/community/:id", controller.CommIntroByIdController)

		// 发布信息
		v1.POST("/post", controller.PostMsgController)
		// 通过帖子id查看帖子
		v1.GET("/post/:id", controller.GetPostMsgByIdController)
		// 查看分页帖子详情posts?page=x&size=x
		v1.GET("/posts", controller.GetPostMsgListController)

		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}
	return r
}
