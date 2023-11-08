package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	DBNAME     = "go_db"
	DBUSER     = "root"
	DBPASSWORD = ""
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASSWORD, DBNAME)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Maʼlumotlar bazasiga ulanib boʻlmadi")
	}

	db.AutoMigrate(&User{})
	seedUsers(db)
}
