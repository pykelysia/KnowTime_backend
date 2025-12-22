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
	return
}
