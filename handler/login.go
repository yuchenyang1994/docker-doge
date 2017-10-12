package handler

import (
	"docker-doge/db"

	"strconv"

	"github.com/gin-gonic/gin"
)

func JwtAuthenticatorHandler(username string, password string, c *gin.Context) (string, bool) {
	d := db.GetDbInstance()
	user := &db.User{Email: username, Password: password}
	user, has := user.GetUserByPassword(d)
	userId := strconv.Itoa(int(user.ID))
	if has {
		return userId, true
	}
	return userId, false
}

func JwtAuthorizatorHandler(userId string, c *gin.Context) bool {
	d := db.GetDbInstance()
	defer d.Close()
	user := &db.User{}
	uuserId, _ := strconv.ParseUint(userId, 0, 64)
	if notFound := d.First(user, uuserId).RecordNotFound(); notFound != true {
		return true
	}
	return false
}

func JwtUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
