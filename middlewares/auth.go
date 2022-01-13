package middlewares

import (
	"log"
	"strings"
	"web_app/controller"
	"web_app/pkg/jwt"

	"github.com/gin-gonic/gin"
)


// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//前端把token放在header的Authorization字段中,值为“Bearer”+token的形式
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			controller.ResponseError(c,controller.CodeNeedLogin )
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken, "Token格式不对")
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		log.Println("token = ",parts[1] )
		if err != nil {
			controller.ResponseError(c,controller.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}