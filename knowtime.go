package main

import (
	"knowtime/config"
	"knowtime/database"
	"knowtime/route"

	"github.com/gin-gonic/gin"
	"github.com/pykelysia/pyketools"
)

func main() {

	// 加载环境
	config.LoadEnv("./.env")

	// 初始化数据库
	database.InitDatabase()

	// 开启网络服务
	server := gin.Default()
	route.Bind(server)         // 绑定路由
	err := server.Run(":8080") // 开始监听
	if err != nil {
		pyketools.Fatalf("net start fail: %v", err)
	}
}
