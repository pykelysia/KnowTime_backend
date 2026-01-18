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
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrUnauthorized, nil))
			return
		}

		var i internal.InternalUsualMsgPostReq
		if err := ctx.ShouldBindJSON(&i); err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrInvalidRequestBody, nil))
			return
		}

		b, err := internal.InternalUsualMsgPostInternal(userIDFromJWT.(uint), i)
		if err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(b.ErrCode, nil))
			return
		}

		ctx.JSON(http.StatusOK, internal.NewResponse(internal.SUCCESS, map[string]any{
			"success": true,
		}))
	}
}
