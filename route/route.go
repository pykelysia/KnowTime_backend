package route

import (
	"knowtime/middleware"
	"knowtime/route/v1/handler"

	_ "knowtime/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token
func Bind(server *gin.Engine) (err error) {
	//服务器状态
	server.GET("/ping", statusHandler())
	//Swagger文档
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//用户相关路由
	userGroup := server.Group("/user")
	{
		userGroup.POST("/login", userLoginHandler())
		userGroup.POST("/logup", userLogupHandler())
		userGroup.GET("/info/:u_id", middleware.JWTAuthMiddleware(), userInfo())
	}
	//核心路由
	routeV1 := server.Group("/v1")
	{
		msgGroup := routeV1.Group("/msg")
		{
			msgGroup.Use(middleware.JWTAuthMiddleware())

			msgGroup.POST("/post", handler.MsgPostHandler())
		}
		reportGroup := routeV1.Group("/report")
		{
			reportGroup.Use(middleware.JWTAuthMiddleware())

			reportGroup.POST("/:date", handler.GenerateHandler())
		}
		chatGroup := routeV1.Group("/chat")
		{
			chatGroup.Use(middleware.JWTAuthMiddleware())

			chatGroup.POST("/chat", handler.ChatHandle())
		}
	}

	return
}
