package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()

	// 8MB
	router.MaxMultipartMemory = 8 << 20

	// set static
	router.Static("/", "./public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "get form err:%s", err.Error())
			return
		}

		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusInternalServerError, "upload file err:%s", err.Error())
				return
			}
		}

		c.String(http.StatusOK, "uploaded successfully %d files with fileds name=%s and email=%s", len(files), name, email)

	})

	router.Run(":8080")
}
