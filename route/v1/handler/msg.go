package handler

import (
	"knowtime/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MsgPostHandler godoc
//
//	@Summary		发布消息
//	@Description	提交一条消息记录
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			u_id	query		int									true	"用户ID"
//	@Param			message	body		internal.InternalUsualMsgPostReq	true	"消息内容"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		401		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Security		BearerAuth
//	@Router			/v1/msg/post [post]
func MsgPostHandler() gin.HandlerFunc {
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
