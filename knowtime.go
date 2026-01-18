package main

import (
	"knowtime/config"
	"knowtime/database"
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

	// 加载环境
	config.LoadEnv("./.env")

	// 初始化数据库
	database.InitDatabase()

	// 开启网络服务
	server := gin.Default()
	route.Bind(server)         // 绑定路由
	err := server.Run(":8080") // 开始监听
	pyketools.Infof("Gin Server Ready at 8080")
	if err != nil {
		pyketools.Fatalf("net start fail: %v", err)
	}
}
