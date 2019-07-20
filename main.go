package main

import (
	"github.com/gin-gonic/gin"
  	"github.com/jinzhu/gorm"
	  _ "github.com/jinzhu/gorm/dialects/sqlite"
)


func main() {	
	router := gin.Default()
	router.Static("/index", "./public")
	router.GET("/", func(c *gin.Context){
		c.Redirect(301, "/index")
	})
	router.GET("/setup", func(c *gin.Context){
		db, err := gorm.Open("sqlite3", "image.db")
		if err != nil {
			panic("can not open")
		}
		defer db.Close()
		db.AutoMigrate(&Image{})
	})	
	router.POST("/upload", Upload)
	router.Run(":8080")
}
