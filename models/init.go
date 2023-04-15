package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:************@tcp(127.0.0.1:3306)/videocut?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panicln("gorm Init Error:", err)
	}
	return db
}
