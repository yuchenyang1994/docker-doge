package main

import (
	"docker-doge/middleware"

	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func main() {
	e := middleware.GetAuthzInstance()
	r := gin.New()
	r.Use(gin.Logger())   // 日志处理
	r.Use(gin.Recovery()) // 500不处理
	r.Use(authz.NewAuthorizer(e))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Golang",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
