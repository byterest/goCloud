package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"
	"github.com/google/uuid"
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
	db.Create(&Article{Title: title, Content: content, UUID: uuid})
	c.Redirect(301, "/")
}
