package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// LoginRequest 登陆信息
// 要使用 binding 功能， 要打上 binding tag
type LoginRequest struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// bindJson binding json demo
// curl -v -X POST 'http://localhost:9090/bind_json' -H 'Content-Type: application/json' -d '{"user": "zyy","password": "123456"}'
func bindJson(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.User != "zyy" || req.Password != "123456" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "welcome"})
}

func main() {
	router := gin.Default()

	router.POST("/bind_json", bindJson)

	err := router.Run(":9090")
	if err != nil {
		log.Fatal("cannot start web server,", err)
	}
}
