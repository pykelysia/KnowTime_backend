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
//	@Router			/v1/report/{date} [get]
func GenerateHandler() gin.HandlerFunc {
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
		iReq.UId = userIDFromJWT.(uint)

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
