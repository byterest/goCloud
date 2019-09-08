package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
	"github.com/gin-contrib/sessions"
)

func GetALL(c *gin.Context) {
	db := OpenDB()
	defer db.Close()
	var images []Image
	db.Order("created_at desc").Find(&images)
	c.HTML(200, "all.tmpl", gin.H{
		"images": images,
	})
}

func Up(c *gin.Context) {
	c.HTML(200, "up.tmpl", gin.H{
		"title": "Main website",
	})
}

func Write(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.Redirect(303, "/")
	}
	c.HTML(200, "write.tmpl", gin.H{
		"title": "Main website",
	})
}

func SetUp(c *gin.Context) {
	db := OpenDB()
	defer db.Close()
	db.AutoMigrate(&Image{}, &Article{}, &User{}, &Topic{})
}

func GetArticle(c *gin.Context) {
	uuid := c.Param("uuid")
	db := OpenDB()
	defer db.Close()
	var article Article
	db.Where("uuid = ?", uuid).First(&article)
	title := article.Title
	content := article.Content
	t := article.CreatedAt
	time := t.Format("2006-01-02")
	var user User
	db.Model(&article).Association("User").Find(&user)
	username := user.UserName
	if username == "" {
		username = "Unkown"
	}
	c.HTML(200, "post.tmpl", gin.H{
		"username": username,
		"title":   title,
		"content": content,
		"time": time,
	})
}

func GetArticles(c *gin.Context) {
	db := OpenDB()
	defer db.Close()
	var articles []Article
	db.Order("created_at desc").Find(&articles)
	c.HTML(200, "allarticle.tmpl", gin.H{
		"articles": articles,
	})
}

func Generate(c *gin.Context) {
	db := OpenDB()
	defer db.Close()
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
	c.HTML(200, "register.tmpl", gin.H{
	})
	}
}

func Login(c *gin.Context)  {
	session := sessions.Default(c)
	user := session.Get("user")
	if user != nil {
		c.Redirect(303, "/")
	}
	c.HTML(200, "login.tmpl", gin.H{
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
	db := OpenDB()
	defer db.Close()
	var user User
	db.Where("user_name=?", v).First(&user)
	var articles []Article
	db.Model(&user).Related(&articles)
	c.HTML(200, "user.tmpl", gin.H{
		"user": user,
		"articles": articles,
	})
	
}

func GetUserInfo(c *gin.Context)  {
	
}

func EditPost(c *gin.Context)  {
	uuid := c.Param("uuid")
	db := OpenDB()
	var article Article
	defer db.Close()
	db.Where("uuid = ?", uuid).First(&article)
	c.HTML(200, "edit.tmpl", gin.H{
		"title": article.Title,
		"content": article.Content,
		"uuid": article.UUID,
	})
}

func AddTopic(c *gin.Context)  {
	c.HTML(200, "addtopic.tmpl", gin.H{
		
	})
}

func GetTopics(c *gin.Context)  {
	db := OpenDB()
	defer db.Close()
	var topics []Topic
	db.Find(&topics)
	c.HTML(200, "topics.tmpl", gin.H{
		"topics": topics,
	})
}

func GetTopic(c *gin.Context)  {
	db := OpenDB()
	defer db.Close()
	var topic Topic
	title := c.Param("topictitle")
	db.Where("title = ?", title).Find(&topic)
	ID := topic.ID
	var articles []Article
	db.Where("topic_id = ?", ID).Find(&articles)
	desc := topic.Description
	logo := topic.ImageURL
	c.HTML(200, "topic.tmpl", gin.H{
		"title": title,
		"desc" : desc,
		"logo" : logo,
		"articles": articles,
	})
}

func AboutPage(c *gin.Context)  {
	c.HTML(200, "about.tmpl", gin.H{

	})
}

func ContactPage(c *gin.Context)  {
	c.HTML(200, "contact.tmpl", gin.H{

	})
}