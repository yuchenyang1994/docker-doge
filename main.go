package main

import (
	"fmt"
	"os"

	"docker-doge/handler/validators"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var useage = `
NAME:
   docker-doge is a usefully docker gui control panel.
   For more details at https://github.com/yuchenyang1994/docker-doge
USAGE:
    run : run the docker-doge server
    migrate: miagate docker-doge database
`

func runServer() {
	r := gin.New()
	validators.RegisterV() // 注册验证器
	URL(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		fmt.Println(arg)
		switch arg {
		case "run":
			runServer()
		case "migrate":
			migrate()
		case "createrole":
			createUserGroup()
		}
	} else {
		fmt.Println(useage)
	}

}
