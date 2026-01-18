package handler

import (
	"knowtime/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GenerateHandler godoc
//
//	@Summary		生成报告
//	@Description	根据日期生成报告
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param			request	body		internal.InternalGenerateReq	true	"请求体"
//	@Success		200		{object}	map[string]interface{}
//	@Failure		400		{object}	map[string]interface{}
//	@Failure		401		{object}	map[string]interface{}
//	@Failure		500		{object}	map[string]interface{}
//	@Security		BearerAuth
//	@Router			/v1/report/{date} [post]
func GenerateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userIDFromJWT, exists := ctx.Get("user_id")
		userIDFromParma, err := strconv.Atoi(ctx.Query("u_id"))
		if !exists || userIDFromJWT.(uint) != uint(userIDFromParma) || err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrUnauthorized, nil))
			return
		}

		var iReq internal.InternalGenerateReq
		if err := ctx.ShouldBindJSON(&iReq); err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrInvalidRequestBody, nil))
			return
		}
		iReq.UId = userIDFromJWT.(uint)

		iResp, b, err := internal.InternalGenerateInternal(iReq)
		if err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(b.ErrCode, nil))
			return
		}

		ctx.JSON(http.StatusOK, internal.NewResponse(internal.SUCCESS, iResp))
	}
}
