package route

import "github.com/gin-gonic/gin"

func Bind(server *gin.Engine) (err error) {
	userGroup := server.Group("/user")
	{
		userGroup.POST("/login", userLoginHandler())
		userGroup.POST("/logup", userLogupHandler())
		userGroup.POST("/info", userInfo())
	}

	internalGroup := server.Group("/internal")
	{
		internalGroup.POST("/usual-msg-post", internalUsualMsgPost())
		internalGroup.POST("/generate", internalGenerate())
	}
	return
}
