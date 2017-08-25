package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	r := gin.New()
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	r.Use(gin.Logger())   // 日志处理
	r.Use(gin.Recovery()) // 500不处理
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Golang",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
