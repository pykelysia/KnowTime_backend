package database

type (
	User struct {
		UId uint `gorm:"primaryKey; autoIncrement; column:uId"`
		// password need use hash to secret
		Password   string      `gorm:"column:password"`
		Name       string      `gorm:"unique; column:name"`
		TimeEvents []TimeEvent `gorm:"foreignKey:UIdRefer"`
	}
	TimeEvent struct {
		TimeEventId uint `gorm:"primaryKey; autoIncrement; column:timeEventId"`
		// format yyyy-mm-dd
		Date    string `gorm:"column:date"`
		AppName string `gorm:"column:appName"`
		// ms
		Duration int  `gorm:"column:duration"`
		UIdRefer uint `gorm:"column:uIdRefer"`
	}
)
