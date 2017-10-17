package handler

import (
	"docker-doge/middleware/policy"

	"docker-doge/db"

	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUserGroupHandler create a UserGroup
func CreateUserGroupHandler(c *gin.Context) {
	groupName := c.Param("groupName")
	d := db.GetDbInstance()
	usergroup := db.UserGroup{GroupName: groupName}
	d.NewRecord(&usergroup)
	d.Create(&usergroup)
	// Add this usergroup policy
	policy.AddPolicyForUserGroups(usergroup.GroupName)
	c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
}

// RemoveUserGroupHandler remove a UserGroup
func RemoveUserGroupHandler(c *gin.Context) {
	groupName := c.Param("groupName")
	d := db.GetDbInstance()
	policy.RemovePolicyForUserGroups(groupName)
	usergroup := db.UserGroup{}
	d = d.First(&usergroup, "group_name = ?", groupName)
	if d.RecordNotFound() != true {
		d.Delete("user_group_id = ?", usergroup.ID)
		d.Delete(db.UserGroup{}, usergroup.ID)
		c.JSON(http.StatusOK, gin.H{"message": "success", "status": http.StatusOK})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "not Found"})
	}

}

// ChangeUserGroupNameHandler ...
func ChangeUserGroupNameHandler(c *gin.Context) {
	groupName := c.Param("groupName")
	d := db.GetDbInstance()
	usergroup := db.UserGroup{}
	d = d.First(&usergroup, "group_name = ?", groupName)
	if d.RecordNotFound() != true {
		policy.RemovePolicyForUserGroups(groupName)
		usergroup.GroupName = groupName
		d.Save(&usergroup)
		policy.AddPolicyForUserGroups(groupName)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "not found"})

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
