package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

var DB *gorm.DB

func SetUp(path string) {
	db, err := gorm.Open("postgres", path)
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)
	// 禁用表名复数
	db.SingularTable(true)
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}

func migration() {
	DB.Set("gorm: table_options", "charset=utf8mb4").
		AutoMigrate(&Apk{})
}
