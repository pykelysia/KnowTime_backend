package handler

import (
	"knowtime/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
