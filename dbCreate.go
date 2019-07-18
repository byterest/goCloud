package main

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Image struct {
	gorm.Model
	ImageUrl string
}
func main()  {
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open")
	}
	defer db.Close()
	db.AutoMigrate(&Image{})
}