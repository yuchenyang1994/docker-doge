package main

import (
	"docker-doge/configs"
	"docker-doge/db"
	"docker-doge/middleware"

	"fmt"
	"os"

	"docker-doge/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// run ...
func runServer() {
	e := middleware.GetAuthzInstance()
	r := gin.New()
	r.Use(gin.Logger())                            // 日志处理
	r.Use(gin.Recovery())                          // 500不处理
	jwtMiddleWare := configs.NewJwtMiddleWare()    // jwt中间件
	jwtAuthorizer := middleware.NewJwtAuthorizer() // jwt权限校验器
	r.POST("/login", jwtMiddleWare.LoginHandler)
	r.POST("/register", handler.RegisterHandler)
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
