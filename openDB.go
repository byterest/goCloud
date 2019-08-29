package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open the db")
	}
	return db
}