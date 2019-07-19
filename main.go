package main

import (
	"fmt"
	"net/http"
	"path"
	"crypto/md5"
	"io"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Image struct {
	gorm.Model
	ImageUrl string
}

func main() {
	unixT := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(unixT,10))
	token := h.Sum(nil)
	router := gin.Default()
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file") // err can handle the file did not exist
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
			db, err := gorm.Open("sqlite3","image.db")
			if err != nil {
			panic("cant not open")
			}
			defer db.Close()
			db.Create(&Image{ImageUrl: filename})


			c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully in images/%s", file.Filename,filename))
	})
	router.Run(":8080")
}
