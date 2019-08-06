package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/ikeikeikeike/go-sitemap-generator.v2/stm"
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

func GetArticle(c *gin.Context) {
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
		"title":   title,
		"content": content,
	})
}

func GetArticles(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open the db")
	}
	var articles []Article
	db.Find(&articles)
	c.HTML(200, "allarticle.html", gin.H{
		"articles": articles,
	})
}

func Generate(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open the db")
	}
	var articles []Article
	db.Find(&articles)
	sm := stm.NewSitemap(1)
	sm.Create()
	sm.SetCompress(false)
	sm.SetPublicPath("public/")
	sm.SetDefaultHost("https://www.prejudice.io")
	sm.Add(stm.URL{{"loc", ""}, {"changefreq", "always"}, {"mobile", true}})
	sm.Add(stm.URL{{"loc", "readme"}})
	sm.Add(stm.URL{{"loc", "aboutme"}, {"priority", 0.1}})
	for _,v := range articles {
			sm.Add(stm.URL{{"loc", "post/"+v.UUID}})
	}
	sm.SetVerbose(true)
	sm.Finalize()
}