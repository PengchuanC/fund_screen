package connector

import (
	"fund_screen/common/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB


func init() {
	var (
		err error
		setting settings.Settings
	)
	setting = settings.GetSettings()
	db, err = gorm.Open(mysql.Open(setting.Source), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败，连接dsn为\n\t%s\n", setting.Source)
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}