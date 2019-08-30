package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
//Image but I want change it to File
type Image struct {
	gorm.Model
	ImageUrl string
	FileName string `gorm:"default:'NoName'"`
}
//Article but I want to change it to Post
type Article struct {
	gorm.Model
	Title     string
	Content   string
	UUID      string
	UserId    int `gorm:"default:0"`
	User      User
	Edited    bool `gorm:"default:false"`
	TopicsID  uint
}
// User 
type User struct {
	gorm.Model
	UserName string
	Password string
}
// Topics
type Topics struct {
	gorm.Model
	Title        string
	ImageURL     string
	Description  string
}
