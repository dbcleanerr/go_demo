package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func index(ctx *gin.Context) {
	name, ok := ctx.Get("name")
	if ok {
		ctx.JSON(http.StatusOK, gin.H{"message": name})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "not found"})
	}

}

func video(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "video"})
	time.Sleep(2 * time.Second)
	return
}

func m1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("name", "zhangyaoyao")
		fmt.Println("m1...")
		start := time.Now()
		ctx.Next()
		cost := time.Since(start)

		fmt.Println(cost)
	}
}

func main() {
	router := gin.Default()

	//router.Group("/").Use(m1())
	router.Use(m1())
	router.GET("index", index)
	router.GET("video", video)

	router.Run(":8080")

}
