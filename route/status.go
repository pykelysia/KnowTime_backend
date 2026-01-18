package route

import (
	"knowtime/database"
	"knowtime/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStatus godoc
//
//	@Summary		获取服务状态
//	@Description	检查Gin和数据库服务的状态
//	@Tags			Status
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}
//	@Failure		500	{object}	map[string]interface{}
//	@Router			/ping [get]
func statusHandler() gin.HandlerFunc {
	// 检查数据库连接
	if err := database.InitDatabase(); err != nil {
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
