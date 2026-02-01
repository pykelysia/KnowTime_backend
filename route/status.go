package route

import (
	"knowtime/database"
	"knowtime/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pykelysia/pyketools"
)

// GetStatus godoc
//
//	@Summary		获取服务状态
//	@Description	检查Gin和数据库服务的状态，始终返回HTTP 200，错误通过响应体中的errcode判断
//	@Tags			Status
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	internal.Response	"服务状态检查结果，errcode=0表示成功"
//	@Router			/ping [get]
func statusHandler() gin.HandlerFunc {
	// 检查数据库连接
	if err := database.InitDatabase(); err != nil {
		pyketools.Infof("DB Failed: %v", err)
		return func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, internal.NewResponse(internal.ErrDatabaseConnection, map[string]any{
				"Gin":      "pong",
				"Database": err.Error(),
			}))
		}
	}
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, internal.NewResponse(internal.SUCCESS, map[string]any{
			"Gin":      "pong",
			"Database": "pong",
		}))
	}
}
