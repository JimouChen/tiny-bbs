package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"tiny-bbs/controller"
	"tiny-bbs/pkg/jwt"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 所以拿到的token格式是：Bearer tokenxxxx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseErr(c, controller.CodeAuthNull)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseErr(c, controller.CodeAuthErrFormat)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseErr(c, controller.CodeAuthInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username或者userId信息保存到请求的上下文c上
		c.Set(controller.KeyUid, mc.UserId)
		c.Set(controller.KeyUname, mc.Username)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
