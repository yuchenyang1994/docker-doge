package main

import (
	"docker-doge/db"
	"docker-doge/middleware"

	"fmt"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

// run ...
func runServer() {
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

// migrate ...
func migrate() {
	d := db.GetDbInstance()
	db.MigrationDB(d)
}

// createRole ...
func createRole() {
	d := db.GetDbInstance()
	db.CreateRole(d)
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		fmt.Println(arg)
		switch arg {
		case "runserver":
			runServer()
		case "migrate":
			migrate()
		case "createrole":
			createRole()
		}
	} else {
		fmt.Println("USAGE:",
			"runserver: 启动服务器",
			"migrate: 同步表结构",
			"createrole: 插入初始用户角色")
	}

}
