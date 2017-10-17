package handler

import (
	"docker-doge/middleware/policy"

	"docker-doge/db"

	"net/http"

	"docker-doge/handler/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateUserGroupHandler create a UserGroup
func CreateUserGroupHandler(c *gin.Context) {
	var vUserGroup validators.UserGroupVlidator
	if err := c.BindWith(&vUserGroup, binding.JSON); err == nil {
		groupName := vUserGroup.GroupName
		d := db.GetDbInstance()
		usergroup := db.UserGroup{GroupName: groupName}
		d.NewRecord(&usergroup)
		d.Create(&usergroup)
		// Add this usergroup policy
		policy.AddPolicyForUserGroups(usergroup.GroupName)
		c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"code":    http.StatusBadRequest,
			"message": "valid fail"})
	}
}

// RemoveUserGroupHandler remove a UserGroup
func RemoveUserGroupHandler(c *gin.Context) {
	var vUserGroup validators.UserGroupVlidator
	if err := c.BindWith(&vUserGroup, binding.JSON); err == nil {
		groupName := vUserGroup.GroupName
		d := db.GetDbInstance()
		policy.RemovePolicyForUserGroups(groupName)
		usergroup := db.UserGroup{}
		d = d.First(&usergroup, "group_name = ?", groupName)
		if !d.RecordNotFound() {
			d.Delete("user_group_id = ?", usergroup.ID)
			d.Delete(db.UserGroup{}, usergroup.ID)
			c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "not Found"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"code":    http.StatusBadRequest,
			"message": "valid fail"})
	}

}

// ChangeUserGroupNameHandler ...
func ChangeUserGroupNameHandler(c *gin.Context) {
	var vUserGroup validators.ChangeGroupNameVlidator

	if err := c.BindWith(&vUserGroup, binding.JSON); err == nil {
		groupName := vUserGroup.GroupName
		newGroupName := vUserGroup.NewGroupName
		d := db.GetDbInstance()
		usergroup := db.UserGroup{}
		d = d.First(&usergroup, "group_name = ?", groupName)
		if !d.RecordNotFound() {
			policy.RemovePolicyForUserGroups(groupName)
			usergroup.GroupName = newGroupName
			d.Save(&usergroup)
			policy.AddPolicyForUserGroups(newGroupName)
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusOK, "message": "susscess"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "not found"})

		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
			"code":    http.StatusBadRequest,
			"message": "valid fail"})
	}

}

// GetUserGroupsHandler ...
func GetUserGroupsHandler(c *gin.Context) {
	usergroup := db.UserGroup{}
	usergroups := usergroup.GetUserGroups()
	if len(usergroups) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"groups": usergroups,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusOK,
			"message": "not found",
		})
	}
}
