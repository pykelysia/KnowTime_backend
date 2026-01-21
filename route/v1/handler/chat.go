package handler

import (
	"knowtime/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatHandle handles chat requests for an authenticated user.
//
// @Summary      Chat with the service
// @Description  Processes a chat request for the authenticated user identified by the JWT and `u_id` query parameter.
// @Tags         chat
// @Accept       json
// @Produce      json
// @Param        u_id   query     int     true  "User ID"
// @Param        request body     internal.InternalChatReq true "Chat request payload"
// @Success      200    {object}  internal.Response "Successful chat response"
// @Failure      200    {object}  internal.Response "Error response with appropriate error code"
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
