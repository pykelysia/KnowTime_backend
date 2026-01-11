package route

import (
	"knowtime/middleware"

	"github.com/gin-gonic/gin"
)

func Bind(server *gin.Engine) (err error) {
	userGroup := server.Group("/user")
	{
		userGroup.POST("/login", userLoginHandler())
		userGroup.POST("/logup", userLogupHandler())
		userGroup.GET("/info", middleware.JWTAuthMiddleware(), userInfo())
	}

	internalGroup := server.Group("/internal")
	internalGroup.Use(middleware.JWTAuthMiddleware())
	{
		internalGroup.POST("/usual-msg-post", internalUsualMsgPostHandler())
		internalGroup.POST("/generate", internalGenerateHandler())
	}

	routeV1 := server.Group("/v1")
	{
		msgGroup := routeV1.Group("/msg")
		{
			msgGroup.Use(middleware.JWTAuthMiddleware())

			msgGroup.POST("/post")
		}
		reportGroup := routeV1.Group("/report")
		{
			reportGroup.Use(middleware.JWTAuthMiddleware())
			reportGroup.GET("/:date")
		}
	}

	return
}
