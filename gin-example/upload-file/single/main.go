package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()

	// 8MB, default 32MB
	router.MaxMultipartMemory = 8 << 20

	// 设置 static，只有设置了的静态目录才能访问
	// / 为访问的根目录
	// ./public为本地的目录
	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// file
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err:%s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusInternalServerError, "upload file err:%s", err.Error())
			return
		}

		c.String(http.StatusOK, "file %s uploaded successfully with fields name=%s and email=%s", file.Filename, name, email)
	})

	router.Run(":8080")
}
