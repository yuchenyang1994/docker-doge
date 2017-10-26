package handler

import "github.com/gin-gonic/gin"
import "docker-doge/handler/validators"
import "github.com/gin-gonic/gin/binding"
import "docker-doge/middleware"
import "strconv"
import "net/http"
import "docker-doge/handler/services"

// AddRoleForUsers ...
func AddRoleForUsers(c *gin.Context) {
	var vAddRole validators.AddRoleValidator
	if err := c.ShouldBindWith(&vAddRole, binding.JSON); err == nil {
		e := middleware.GetAuthzInstance(c)
		ok := e.AddRoleForUser(strconv.Itoa(int(vAddRole.UserID)), vAddRole.RoleName)
		if ok {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusBadRequest, "message": "role is existed"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "valid fail", "error": err.Error()})
	}
}

// RemoveRoleForUsers ....
func RemoveRoleForUsers(c *gin.Context) {
	var vAddRole validators.AddRoleValidator
	if err := c.ShouldBindWith(&vAddRole, binding.JSON); err == nil {
		e := middleware.GetAuthzInstance(c)
		ok := e.DeleteRoleForUser(strconv.Itoa(int(vAddRole.UserID)), vAddRole.RoleName)
		if ok {
			c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "success"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "role is not existed"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "valid fail", "error": err.Error()})
	}

}

// GetUsersInfos ...
func GetUsersInfos(c *gin.Context) {
	groupName := c.Param("groupName")
	e := middleware.GetAuthzInstance(c)
	userinfoService := services.NewUserInfoService(e, groupName, c)
	userinfos := userinfoService.GetUserInfos()
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": userinfos})
}
