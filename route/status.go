package route

import (
	"knowtime/database"

	"github.com/gin-gonic/gin"
)

// statusHandler 返回服务状态信息
func statusHandler() gin.HandlerFunc {
	// 检查数据库连接
	if err := database.InitDatabase(); err != nil {
		return func(ctx *gin.Context) {
			ctx.JSON(500, gin.H{
				"status":   "error",
				"Gin":      "pong",
				"Database": err.Error(),
			})
		}
	}
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":   "ok",
			"Gin":      "pong",
			"Database": "pong",
		})
	}
}
