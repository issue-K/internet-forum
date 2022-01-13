package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"

	_ "web_app/docs"  // 千万不要忘了导入把你上一步生成的docs

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup()( *gin.Engine ){
	r := gin.New()
	r.Use( logger.GinLogger(),logger.GinRecovery(true) )
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static","./static")
	r.GET("/",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)

	v1.POST("/login", controller.LoginHandler)


	v1.Use(middlewares.JWTAuthMiddleware() )
	{
		v1.GET("/community",controller.CommunityHandler)
		v1.GET("/",middlewares.JWTAuthMiddleware(),func(c *gin.Context){
			c.String( http.StatusOK,"ok")
		})
		v1.GET("/community/:id",controller.CommunityDetailHandler)

		v1.POST("/post",controller.PostCommitHandler )
		v1.GET("/post/:id",controller.GetPostDetailHandler )
		v1.GET("/posts2",controller.PostListHandler)

		v1.POST("/vote",controller.POSTVoteController )
	}
	return r
}
