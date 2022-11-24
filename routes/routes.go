package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	gs "github.com/swaggo/gin-swagger"
	"time"

	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-contrib/cors"
	"github.com/swaggo/files"
	"net/http"
	"tiny-bbs/controller"
	_ "tiny-bbs/docs" // 千万不要忘了导入把你上一步生成的docs
	"tiny-bbs/middleware"
)

func Init() *gin.Engine {
	//r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", middleware.RateLimitMiddleware(time.Second*2, 1), func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"version": viper.GetString("app.version"),
		})
	})

	v1 := r.Group("/api/v1")

	// 用户注册
	v1.POST("/signup", controller.SignUpController)
	// 用户登陆
	v1.POST("/login", controller.LoginController)
	v1.GET("/community", controller.CommunityController)
	v1.GET("/community/:id", controller.CommIntroByIdController)

	// 通过帖子id查看帖子
	v1.GET("/post/:id", controller.GetPostMsgByIdController)
	// 查看分页帖子详情posts?page=x&size=x
	v1.GET("/posts", controller.GetPostMsgListController)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		// 发布信息
		v1.POST("/post", controller.PostMsgController)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	return r
}
