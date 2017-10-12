package handler

import (
	"docker-doge/db"
	"docker-doge/handler/validators"

	"docker-doge/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

func RegisterHandler(c *gin.Context) {
	var singin validators.SingupVlidator
	if err := c.ShouldBindWith(&singin, binding.JSON); err == nil {
		d := db.GetDbInstance()
		if ok := registerUser(singin.Email, singin.Password, singin.UserGroupID, d); ok {
			c.JSON(http.StatusOK, gin.H{"message": "sussess", "status": http.StatusOK})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "fail register", "status": http.StatusBadRequest})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "valid fail", "error": err.Error()})
	}
}

func registerUser(email string, password string, groupID uint, d *gorm.DB) bool {
	user := &db.User{Email: email, Password: password, UserGroupID: groupID}
	if err := user.Insert(d); err == nil {
		e := middleware.GetAuthzInstance()
		// 注册用户角色
		if ok := e.AddRoleForUser(user.Email, middleware.ROLE_USER); ok {
			return true
		}
		return false
	}
	return false
}
