package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func internalUsualMsgPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, BaseMsg{
				Code:    401,
				Message: "User not authenticated",
			})
			return
		}

		// TODO: 实现具体业务逻辑，这里可以使用userID

		ctx.JSON(http.StatusOK, BaseMsg{
			Code:    200,
			Message: "User ID: " + string(rune(userID.(uint))),
		})
	}
}

func internalGenerate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, BaseMsg{
				Code:    401,
				Message: "User not authenticated",
			})
			return
		}

		// TODO: 实现具体业务逻辑，这里可以使用userID

		ctx.JSON(http.StatusOK, BaseMsg{
			Code:    200,
			Message: "Authenticated user ID: " + string(rune(userID.(uint))),
		})
	}
}
