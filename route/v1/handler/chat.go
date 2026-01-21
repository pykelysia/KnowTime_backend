package handler

import (
	"knowtime/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatHandler godoc
//
//	@Summary		与AI对话
//	@Description	对于登录用户，使用此接口与AI进行对话，始终返回HTTP 200，错误通过响应体中的errcode判断
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			request	body		internal.InternalChatReq	true	"对话上下文请求体"
//	@Success		200		{object}	internal.Response			"对话返回响应，errcode=0表示成功"
//	@Security		BearerAuth
//	@Router			/v1/chat [post]
func ChatHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var iReq internal.InternalChatReq
		if err := ctx.ShouldBindJSON(&iReq); err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrInvalidRequestBody, nil))
			return
		}

		userIDFromJWT, exists := ctx.Get("user_id")
		userIDFromParma := iReq.UId
		if !exists || userIDFromJWT.(uint) != uint(userIDFromParma) {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrUnauthorized, nil))
			return
		}

		iResp, b, err := internal.InternalChatInternal(iReq)
		if err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(b.ErrCode, nil))
			return
		}

		ctx.JSON(http.StatusOK, internal.NewResponse(internal.SUCCESS, iResp))
	}
}
