package handler

import (
	"docker-doge/db"
	"docker-doge/handler/validators"

	"docker-doge/middleware"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

func RegisterHandler(c *gin.Context) {
	var singin validators.SingupVlidator
	if err := c.ShouldBindWith(&singin, binding.JSON); err == nil {
		d := db.GetDbInstance(c)
		if ok := registerUser(singin.Email, singin.Password, singin.UserGroupID, d, c); ok {
			c.JSON(http.StatusOK, gin.H{"message": "sussess", "status": http.StatusOK})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "fail register", "status": http.StatusBadRequest})
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "valid fail", "error": err.Error()})
	}
}

func registerUser(email string, password string, groupID uint, d *gorm.DB, c *gin.Context) bool {
	user := &db.User{Email: email, Password: password, UserGroupID: groupID}
	if err := user.Insert(d); err == nil {
		e := middleware.GetAuthzInstance(c)
		// 注册用户角色
		strUserId := strconv.Itoa(int(user.ID))
		if ok := e.AddRoleForUser(strUserId, middleware.ROLE_USER); ok {
			return true
		}
		return false
	}
	return false
}
