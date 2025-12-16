package main

import (
	"knowtime/config"
	"knowtime/database"
)

func main() {
	config.LoadEnv("./.env")
	database.InitDatabase()
}
