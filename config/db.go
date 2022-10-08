package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// 初始化db
func InitDB() *gorm.DB {
	//连接到progresql数据库
	db, err := gorm.Open("postgres", "host=192.168.127.128 port=5432 user=postgres dbname=photograph password=123456 sslmode=disable")
	if err != nil {
		// logger.PanicError(err, "链接数据库错误", true)
		panic("failed to connect database, err:" + err.Error())
	}

	//db.AutoMigrate(&model.Users{}) // 自动创建 Users 表

	db.LogMode(true)
	DB = db
	return db
}

// 获取db句柄
func GetDB() *gorm.DB {
	return DB
}
