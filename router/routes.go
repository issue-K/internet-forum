package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
)

func Setup()( *gin.Engine ){
	r := gin.New()
	r.Use( logger.GinLogger(),logger.GinRecovery(true) )

	r.POST("/signup",controller.SignUpHandler )

	r.POST("/login",controller.LoginHandler )

	r.GET("/",func(c *gin.Context){
		c.String( http.StatusOK,"ok")
	})
	return r
}
