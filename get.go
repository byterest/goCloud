package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/ikeikeikeike/go-sitemap-generator.v2/stm"
	"github.com/gin-contrib/sessions"
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
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(303, "/")
	}
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
	db.AutoMigrate(&Image{}, &Article{}, &User{})
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
	var user User
	db.Model(&article).Association("User").Find(&user)
	username := user.UserName
	if username == "" {
		username = "Unkown"
	}
	c.HTML(200, "post.html", gin.H{
		"username": username,
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
	db.Order("created_at desc").Find(&articles)
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

func Signup(c *gin.Context)  {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		c.Redirect(303, "/")
	} else{
	c.HTML(200, "register.html", gin.H{
	})
	}
}

func Login(c *gin.Context)  {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		c.Redirect(303, "/")
	}
	c.HTML(200, "login.html", gin.H{
	})
}

func Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
}

func GetUser(c *gin.Context)  {
	session := sessions.Default(c)
	v := session.Get("user")
	if v == nil {
		c.Redirect(303, "/")
	}
	db, err := gorm.Open("sqlite3", "image.db")
	var user User
	db.Where("user_name=?", v).First(&user)
	var articles []Article
	db.Model(&user).Related(&articles)
	if err != nil {
		panic("db open failed")
	}
	c.HTML(200, "user.html", gin.H{
		"user": user,
		"articles": articles,
	})
	
}

func GetUserInfo(c *gin.Context)  {
	
}