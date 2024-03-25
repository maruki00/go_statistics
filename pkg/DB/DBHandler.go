package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	db *gorm.DB
}

func (obj *DB) GetDB() *gorm.DB {
	dsn := "root:user@tcp(127.0.0.1:3306)/go_appstatics?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Error")
	}

	obj.db = db
	return obj.db
}
