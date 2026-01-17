package route

import (
	"knowtime/middleware"
	"knowtime/route/v1/handler"

	"github.com/gin-gonic/gin"
)

func Bind(server *gin.Engine) (err error) {
	server.GET("/ping", statusHandler())

	userGroup := server.Group("/user")
	{
		userGroup.POST("/login", userLoginHandler())
		userGroup.POST("/logup", userLogupHandler())
		userGroup.GET("/info/:u_id", middleware.JWTAuthMiddleware(), userInfo())
	}

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

			reportGroup.GET("/:date", handler.GenerateHandler())
		}
	}

	return
}
