package main

import (
	"knowtime/config"
	"knowtime/database"
	"knowtime/internal"
	"knowtime/route"

	"github.com/gin-gonic/gin"
	"github.com/pykelysia/pyketools"
)

// @title						知时后端API文档
// @version					1.0
// @description				提供所有API接口的说明文档
// @contact.name				吕舒君
// @contact.email				Lvshujun0918@163.com
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and JWT token
func main() {
	// 打印构建信息
	internal.PrintBuildInfo()
	// 加载环境
	config.LoadEnv("./.env")
	pyketools.Infof("Env Ready!")
	// 初始化数据库
	database.InitDatabase()
	pyketools.Infof("DB Ready!")
	// 开启网络服务
	server := gin.Default()
	// 设置Release模式
	gin.SetMode(gin.ReleaseMode)
	// 绑定路由
	route.Bind(server)
	pyketools.Infof("Route Ready!")
	// 开始监听
	err := server.Run(":8080")
	pyketools.Infof("Gin Server[8080] Ready!")
	if err != nil {
		pyketools.Fatalf("Gin Server[8080] Failed: %v", err)
	}
}
