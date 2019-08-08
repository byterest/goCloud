package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-contrib/sessions"
)

func Upload(c *gin.Context) {
	unixT := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(unixT, 10))
	token := h.Sum(nil)

	file, err := c.FormFile("file") // err can handle the file did not exist
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	newFileName := c.PostForm("filename")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	ext := path.Ext(file.Filename)
	filename := hex.EncodeToString(token) + ext
	if err := c.SaveUploadedFile(file, "public/images/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	// this block I want to insert a imageurl to the table
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("cant not open")
	}
	defer db.Close()
	db.Create(&Image{ImageUrl: filename, FileName: newFileName})
	c.Redirect(301, "/")
}

func WriteIn(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	uuid := uuid.New().String()
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("cant not open")
	}
	defer db.Close()
	session := sessions.Default(c)
	username := session.Get("user")
	if username != nil {
		c.Redirect(303, "/")
	} 
	var user User
	db.Where("user_name = ?", username).First(&user)
	db.Create(&Article{Title: title, Content: content, UUID: uuid, User: user})
	c.Redirect(301, "/")
}

func SignupPost(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password1 := c.PostForm("password1")
	if password != password1 {
		panic("cant not open")
	}
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open")
	}
	defer db.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	password = string(hash)
	db.Create(&User{UserName: username, Password: password})
	c.Redirect(303, "/")
}

func LoginPost(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	db, err := gorm.Open("sqlite3", "image.db")
	if err != nil {
		panic("can not open")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	password2 := string(hash)
	if err != nil {
		panic("can not create")
	}
	var user User
	db.Where("user_name = ?", username).First(&user)
	err = bcrypt.CompareHashAndPassword([]byte(password2), []byte(password))
	session := sessions.Default(c)
	if err != nil {
		panic("pass")
	}
	
	session.Set("user", username)
	session.Save()
	c.Redirect(303, "/user")
}