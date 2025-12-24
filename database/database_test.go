package database

import (
	"knowtime/config"
	"os"
	"testing"

	"github.com/pykelysia/pyketools"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	config.LoadEnv("../.env")
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestDatabaseInit(t *testing.T) {
	err := InitDatabase()
	if err != nil {
		pyketools.Fatalf("database open error: %v", err)
	}

	if db == nil {
		pyketools.Fatalf("database open return nil")
	}
}

func TestUserCRUD(t *testing.T) {
	//初始化数据库
	err := InitDatabase()
	if err != nil {
		pyketools.Fatalf("database open error: %v", err)
	}

	//创建用户
	user := User{
		Name:     "test_user",
		Password: "test_password",
	}
	if err := db.Create(&user).Error; err != nil {
		pyketools.Fatalf("user creat error: %v", err)
	}

	// 查询用户
	var foundUser User
	if err := db.First(&foundUser, user.UId).Error; err != nil {
		pyketools.Fatalf("user find error: %v", err)
	}

	if foundUser.Name != user.Name {
		pyketools.Fatalf("user find should be %s, but %s", user.Name, foundUser.Name)
	}

	// 更新用户
	newName := "updated_user"
	if err := db.Model(&foundUser).Update("Name", newName).Error; err != nil {
		pyketools.Fatalf("user update error: %v", err)
	}

	// 验证更新结果
	var updatedUser User
	if err := db.First(&updatedUser, user.UId).Error; err != nil {
		pyketools.Fatalf("user updated find error:: %v", err)
	}

	if updatedUser.Name != newName {
		pyketools.Fatalf("user updated find should be %s, but %s", newName, updatedUser.Name)
	}

	// 删除用户
	if err := db.Delete(&updatedUser).Error; err != nil {
		pyketools.Fatalf("user delete error: %v", err)
	}

	// 验证删除结果
	var deletedUser User
	if err := db.First(&deletedUser, user.UId).Error; err != gorm.ErrRecordNotFound {
		pyketools.Fatalf("user deleted error: %v", err)
	}
}

func TestTimeEventCRUD(t *testing.T) {
	// 初始化数据库
	err := InitDatabase()
	if err != nil {
		pyketools.Fatalf("database open error: %v", err)
	}

	// 创建一个用户用于关联时间事件
	user := User{
		Name:     "te_user",
		Password: "te_password",
	}

	if err := db.Create(&user).Error; err != nil {
		pyketools.Fatalf("user creat error: %v", err)
	}

	// 创建时间事件
	timeEvent := TimeEvent{
		Date:     "2025-12-25",
		AppName:  "test_app",
		Duration: 3600000, // 1小时（毫秒）
		UIdRefer: user.UId,
	}

	if err := db.Create(&timeEvent).Error; err != nil {
		pyketools.Fatalf("timeevent create error: %v", err)
	}

	// 查询时间事件
	var foundTimeEvent TimeEvent
	if err := db.First(&foundTimeEvent, timeEvent.TimeEventId).Error; err != nil {
		pyketools.Fatalf("timeevent find error: %v", err)
	}

	if foundTimeEvent.AppName != timeEvent.AppName {
		pyketools.Fatalf("timeevent find should be %s, but %s", timeEvent.AppName, foundTimeEvent.AppName)
	}

	// 更新时间事件
	newAppName := "updated_app"
	if err := db.Model(&foundTimeEvent).Update("AppName", newAppName).Error; err != nil {
		pyketools.Fatalf("timeevent update error: %v", err)
	}

	// 验证更新结果
	var updatedTimeEvent TimeEvent
	if err := db.First(&updatedTimeEvent, timeEvent.TimeEventId).Error; err != nil {
		pyketools.Fatalf("timeevent updated find error: %v", err)
	}

	if updatedTimeEvent.AppName != newAppName {
		pyketools.Fatalf("timeevent updated find should be %s, but %s", newAppName, updatedTimeEvent.AppName)
	}

	// 删除时间事件
	if err := db.Delete(&updatedTimeEvent).Error; err != nil {
		pyketools.Fatalf("timeevent delete error: %v", err)
	}

	// 验证删除结果
	var deletedTimeEvent TimeEvent
	if err := db.First(&deletedTimeEvent, timeEvent.TimeEventId).Error; err != gorm.ErrRecordNotFound {
		pyketools.Fatalf("timeevent deleted error: %v", err)
	}

	// 清理用户
	if err := db.Delete(&user).Error; err != nil {
		pyketools.Fatalf("user delete error: %v", err)
	}
}
