package handler

import (
	"docker-doge/db"
	"github.com/gin-gonic/gin"
)

func JwtAuthenticatorHandler(email string, password string, c *gin.Context) (string, bool) {
	d := db.GetDbInstance()
	user := &db.User{Email: email, Password: password}
	user, has := user.GetUserByPassword(d)
	if has == true {
		return email, true
	}
	return email, false
}

func JwtAuthorizatorHandler(email string, c *gin.Context) bool {
	d := db.GetDbInstance()
	defer d.Close()
	user := &db.User{}
	d.Where("email=?", email).First(user)
	if user.Email == email {
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
