package main

import (
	"github.com/gin-gonic/gin"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	router := gin.Default()
	router.Static("/media", "./public/")
	router.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	// Get Method
	router.GET("/files/upload", Up)
	router.GET("/p/new", Write)
	router.GET("/setup", SetUp)
	router.GET("/", GetArticles)
	router.GET("/resources", GetALL)
	router.GET("/post/:uuid", GetArticle)
	router.GET("/sitemapg", Generate)
	router.GET("/signup", Signup)
	router.GET("/login", Login)
	router.GET("/logout", Logout)
	router.GET("/user", GetUser)
	// Post Method
	router.POST("/signup", SignupPost)
	router.POST("/upload", Upload)
	router.POST("/writein", WriteIn)
	router.POST("/login", LoginPost)
	router.Run(":8080")
}
