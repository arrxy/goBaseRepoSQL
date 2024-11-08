package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:root@tcp(127.0.0.1:3307)/book?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %v", err))
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
