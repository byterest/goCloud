package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gin-gonic/gin"
)

func getALL(c *gin.Context){
	db, err := gorm.Open("sqlite3", "image.db")
		if err != nil {
			panic("can not open")
		}
		defer db.Close()
		var images []Image	
		db.Find(&images)
		c.HTML(200, "all.html", gin.H{
			"images": images,
		})
}

func up(c *gin.Context){
	c.HTML(200, "index.html", gin.H{
		"title": "Main website",
	})
}