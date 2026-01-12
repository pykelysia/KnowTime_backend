package route

import (
	"knowtime/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func internalUsualMsgPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDFromJWT, exists := ctx.Get("user_id")
		userIDFromParma, err := strconv.Atoi(ctx.Query("u_id"))
		if !exists || userIDFromJWT.(uint) != uint(userIDFromParma) || err != nil {
			ctx.JSON(http.StatusUnauthorized, BaseMsg{
				Code:    401,
				Message: "User not authenticated",
			})
			return
		}

		var i internal.InternalUsualMsgPostReq
		if err := ctx.ShouldBindJSON(&i); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"net_message": BaseMsg{
					Code:    400,
					Message: "Invalid request body",
				},
			})
			return
		}

		b, err := internal.InternalUsualMsgPostInternal(userIDFromJWT.(uint), i)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"net_message": b,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"net_message": b,
			"data": map[string]any{
				"success": true,
			},
		})
	}
}

func internalGenerateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDFromJWT, exists := ctx.Get("user_id")
		userIDFromParma, err := strconv.Atoi(ctx.Query("u_id"))
		if !exists || userIDFromJWT.(uint) != uint(userIDFromParma) || err != nil {
			ctx.JSON(http.StatusUnauthorized, BaseMsg{
				Code:    401,
				Message: "User not authenticated",
			})
			return
		}

		var iReq internal.InternalGenerateReq
		if err := ctx.ShouldBindJSON(&iReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"net_message": BaseMsg{
					Code:    400,
					Message: "Invalid request body",
				},
			})
			return
		}

		iResp, b, err := internal.InternalGenerateInternal(iReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"net_message": b,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"net_message": b,
			"data":        iResp,
		})
	}
}
