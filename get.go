package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetALL(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open")
	}
	defer db.Close()
	var images []Image
	db.Order("created_at desc").Find(&images)
	c.HTML(200, "all.html", gin.H{
		"images": images,
	})
}

func Up(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Main website",
	})
}

func Write(c *gin.Context) {
	c.HTML(200, "write.html", gin.H{
		"title": "Main website",
	})
}

func SetUp(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open")
	}
	defer db.Close()
	db.AutoMigrate(&Image{}, &Article{})
}

func GetArticle(c *gin.Context)  {
	uuid := c.Param("uuid")
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open")
	}
	defer db.Close()
	var article Article
	db.Where("uuid = ?", uuid).First(&article)
	title := article.Title
	content := article.Content
	c.HTML(200, "post.html", gin.H{
		"title": title,
		"content": content,
	})
}