package handler

import (
	"docker-doge/db"

	"github.com/gin-gonic/gin"
)

func JwtAuthenticatorHandler(username string, password string, c *gin.Context) (string, bool) {
	d := db.GetDbInstance()
	user := &db.User{Email: username, Password: password}
	user, has := user.GetUserByPassword(d)
	if has {
		return user.Email, true
	}
	return user.Email, false
}

func JwtAuthorizatorHandler(userId string, c *gin.Context) bool {
	d := db.GetDbInstance()
	defer d.Close()
	user := &db.User{}
	if notFound := d.First(user, "email = ?", userId).RecordNotFound(); notFound != true {
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
