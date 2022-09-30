package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledata" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"` // CheckOut > CheckIn
}

func main() {
	router := gin.Default()

	// 添加 validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		// tag: 结构体 binding 里指定的
		// fn: 对应的函数
		v.RegisterValidation("bookabledata", bookableDate)
	}

	router.GET("/bookable", func(c *gin.Context) {
		var req Booking
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid"})
	})

	router.Run(":8080")
}

func bookableDate(fl validator.FieldLevel) bool {
	// 根据实际的数据类型指定
	// ok 永远都是 true
	date, ok := fl.Field().Interface().(time.Time)

	fmt.Println(ok)

	if ok {
		today := time.Now()
		// 日期要大于当前日期
		if today.After(date) {
			return false
		}
	}

	return true
}
