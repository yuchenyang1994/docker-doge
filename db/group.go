package db

import (
	"github.com/jinzhu/gorm"
)

// UserGroup Db
type UserGroup struct {
	gorm.Model
	GroupName string `gorm:"size:60"`
	Users     []User
}

// UserGroupById Json format struct
type UserGroupById struct {
	GroupID   uint   `json:"groupId"`
	GroupName string `json:"groupName"`
}

// GetUserGroups return list for Json format struct
func (usergroup UserGroup) GetUserGroups(db *gorm.DB) []UserGroupById {
	groups := []UserGroup{}
	db.Find(&groups)
	groupByIds := []UserGroupById{}
	for _, group := range groups {
		if group.GroupName != "ROOT" {
			groupByid := UserGroupById{GroupID: group.ID, GroupName: group.GroupName}
			groupByIds = append(groupByIds, groupByid)
		}
	}
	return groupByIds

}

// GetUserGroupByName balabala
func (usergroup *UserGroup) GetUserGroupByName(db *gorm.DB) (UserGroupById, bool) {
	db.First(usergroup, "group_name = ?", usergroup.GroupName)
	userGroupbyId := UserGroupById{GroupName: usergroup.GroupName, GroupID: usergroup.ID}
	if db.RecordNotFound() == true {
		return userGroupbyId, false
	}
	return userGroupbyId, true
}

// GetUserGroupById  balabala
func (usergroup *UserGroup) GetUserGroupById(db *gorm.DB) (UserGroupById, bool) {
	db.First(usergroup, usergroup.ID)
	if db.RecordNotFound() == true {
		return UserGroupById{GroupID: usergroup.ID, GroupName: usergroup.GroupName}, false
	}
	return UserGroupById{GroupID: usergroup.ID, GroupName: usergroup.GroupName}, true
}
