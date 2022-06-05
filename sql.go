package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "root:superman@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println(err)
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		Number{},
	)
}
