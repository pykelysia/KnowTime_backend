package main

import (
	"knowtime/config"
	"knowtime/database"
	"knowtime/route"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv("./.env")
	database.InitDatabase()
	server := gin.Default()
	route.Bind(server)
	server.Run(":8080")
}
