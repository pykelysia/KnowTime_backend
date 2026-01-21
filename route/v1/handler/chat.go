package handler

import (
	"knowtime/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChatHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDFromJWT, exists := ctx.Get("user_id")
		userIDFromParma, err := strconv.Atoi(ctx.Query("u_id"))
		if !exists || userIDFromJWT.(uint) != uint(userIDFromParma) || err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrUnauthorized, nil))
			return
		}

		var iReq internal.InternalChatReq
		if err := ctx.ShouldBindJSON(&iReq); err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrInvalidRequestBody, nil))
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
