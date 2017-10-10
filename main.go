package main

import (
	"docker-doge/db"
	"docker-doge/middleware"

	"fmt"
	"os"

	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	jwtMiddleWare := middleware.NewJwtMiddleWare()
	r.POST("/login", jwtMiddleWare.LoginHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}

// migrate ...
func migrate() {
	d := db.GetDbInstance()
	db.MigrationDB(d)
}

// createRole ...
func createUserGroup() {
	d := db.GetDbInstance()
	db.CreateUserGroup(d)
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
			createUserGroup()
		}
	} else {
		fmt.Println("USAGE:",
			"runserver: 启动服务器",
			"migrate: 同步表结构",
			"createrole: 插入用户组角色")
	}

}
