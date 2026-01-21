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
//	@Description	根据日期生成报告，始终返回HTTP 200，错误通过响应体中的errcode判断
//	@Tags			Report
//	@Accept			json
//	@Produce		json
//	@Param			u_id	query		int					true	"用户ID"
//	@Param			date	path		string				true	"日期 (格式: yyyy-mm-dd)"
//	@Success		200		{object}	internal.Response	"报告生成响应，errcode=0表示成功"
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
		iReq.UId = userIDFromJWT.(uint)
		iReq.Date = ctx.Param("date")

		iResp, b, err := internal.InternalGenerateInternal(iReq)
		if err != nil {
			ctx.JSON(http.StatusOK, internal.NewResponse(b.ErrCode, nil))
			return
		}

		ctx.JSON(http.StatusOK, internal.NewResponse(internal.SUCCESS, iResp))
	}
}
