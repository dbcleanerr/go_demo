package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

// upload 上传单个文件
// curl -X POST http://localhost:9090/upload -F "file=@d:\a.txt" -H "Content-Type: multipart/form-data"
func upload(ctx *gin.Context) {
	f, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save
	dst := path.Join("upload_dir", f.Filename)
	err = ctx.SaveUploadedFile(f, dst)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"file_name": f.Filename,
	})
}

// uploadMult 上传多个文件
// curl -X POST http://localhost:9090/uploads -F "upload[]=@d:\a.txt" -F "upload[]=@d:\b.txt" -H "Content-Type: multipart/form-data"
func uploadMult(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["upload[]"]

	for _, f := range files {
		dst := path.Join("upload_dir", f.Filename)
		_ = ctx.SaveUploadedFile(f, dst)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total upload files": len(files),
	})
}

// upload file
func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 16 << 20 // 16MB

	router.POST("/upload", upload)
	router.POST("/uploads", uploadMult)

	err := router.Run(":9090")
	if err != nil {
		log.Fatal("cannot start web server,", err)
	}
}
