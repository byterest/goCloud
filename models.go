package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Image struct {
	gorm.Model
	ImageUrl string
	FileName string `gorm:"default:'NoName'"`
}

type Article struct {
	gorm.Model
	Title   string
	Content string
	UUID    string
	UserId int `gorm:"default:0"`
	User User
}

type User struct {
	gorm.Model
	UserName string
	Password string
}