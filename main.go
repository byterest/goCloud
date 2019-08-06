package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/media", "./public/")
	router.LoadHTMLGlob("templates/*")
	router.GET("/up", Up)
	router.GET("/write", Write)
	router.GET("/setup", SetUp)
	router.GET("/", GetArticles)
	router.GET("/resources", GetALL)
	router.GET("/post/:uuid", GetArticle)
	router.GET("/sitemapg", Generate)
	router.POST("/upload", Upload)
	router.POST("/writein", WriteIn)
	router.Run(":8080")
}
