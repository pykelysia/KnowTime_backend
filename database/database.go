package database

import (
	"fmt"
	"os"

	"github.com/pykelysia/pyketools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (err error) {
	//读取环境变量
	config := DatabaseConfig{
		Username: getEnvWithDefault("DB_USERNAME", "root"),
		Password: getEnvWithDefault("DB_PASSWORD", ""),
		Host:     getEnvWithDefault("DB_HOST", "localhost"),
		Port:     getEnvWithDefault("DB_PORT", "3306"),
		Name:     getEnvWithDefault("DB_NAME", "knowtime"),
	}
	//拼合dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	//连接
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		pyketools.Fatalf("database open error: %v", err)
	}
	//自动迁移
	err = db.AutoMigrate(&User{}, &TimeEvent{})
	if err != nil {
		pyketools.Fatalf("database migrate error: %v", err)
	}
	return
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
