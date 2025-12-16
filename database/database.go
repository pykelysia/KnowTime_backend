package database

import (
	"fmt"
	"os"

	"github.com/pykelysia/pyketools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (err error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbIp := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbIp, dbPort, dbName)
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		pyketools.Fatalf("database open error: %v", err)
	}
	return
}
