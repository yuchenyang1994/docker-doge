package handler

import (
	"docker-doge/handler/validators"

	"docker-doge/db"

	"net/http"

	"fmt"

	"docker-doge/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateUserGroupHandler(c *gin.Context) {
	var userGroupVlidators validators.UserGroupVlidator
	if err := c.ShouldBindWith(&userGroupVlidators, binding.JSON); err == nil {
		d := db.GetDbInstance()
		usergroup := db.UserGroup{GroupName: userGroupVlidators.GroupName}
		d.Create(usergroup)
		// Add this usergroup policy
		AddPolicyForUserGroups(usergroup.GroupName)
		c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "valid fail", "error": err.Error()})
	}
}

func RemoveUserGroupHandler(c *gin.Context) {
	var userGroupIdVlidator validators.UserGroupIdVlidator
	if err := c.ShouldBindWith(&userGroupIdVlidator, binding.JSON); err == nil {
		d := db.GetDbInstance()
		d.Delete(db.UserGroup{}, userGroupIdVlidator.GroupId)
	}
}

func GetUserGroupsHandler(c *gin.Context) {
	usergroup := db.UserGroup{}
	usergroups := usergroup.GetUserGroups()
	c.SecureJSON(http.StatusOK, usergroups)
}

func AddPolicyForUserGroups(groupName string) {
	e := middleware.GetAuthzInstance()
	containerDomin := fmt.Sprintf("/contarner/%s*", groupName)
	// ADMIN Policy
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "GET")
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "POST")
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "PUT")
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "DELETE")
	// Leader Policy
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "GET")
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "POST")
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "PUT")
	// Default Policy
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "GET")
	e.AddPolicy(groupName, middleware.ROLE_ADMIN, containerDomin, "POST")
}
