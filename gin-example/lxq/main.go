package main

import (
	"github.com/dchest/validator"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func parseIP(c *gin.Context) {
	ip := c.Param("ip")
	if net.ParseIP(ip) == nil {
		c.JSON(http.StatusOK, gin.H{"status:": "无效", "IP": ip})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status:": "有效", "IP": ip})
		return
	}
}

func parseDomain(c *gin.Context) {
	domain := c.Param("domain")
	if validator.IsValidDomain(domain) {
		c.JSON(http.StatusOK, gin.H{"status:": "有效", "domain": domain})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status:": "无效", "domain": domain})
		return
	}
}

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.GET("/parse_ip/:ip", parseIP)
	router.GET("/parse_domain/:domain", parseDomain)

	err := router.Run(":2204")
	if err != nil {
		log.Fatal("cannot start web server,err:", err.Error())
	}
}
