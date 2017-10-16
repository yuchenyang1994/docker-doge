package db

import (
	"github.com/jinzhu/gorm"
)

// 用户组
type UserGroup struct {
	gorm.Model
	GroupName string `gorm:"size:60"`
	Users     []User
}

type UserGroupById struct {
	GroupID   uint   `json:"groupId"`
	GroupName string `json:"groupName"`
}

func (usergroup UserGroup) GetUserGroups() []UserGroupById {
	db := GetDbInstance()
	groups := []UserGroup{}
	db.Find(&groups)
	groupByIds := []UserGroupById{}
	for _, group := range groups {
		groupByid := UserGroupById{GroupID: group.ID, GroupName: group.GroupName}
		groupByIds = append(groupByIds, groupByid)
	}
	return groupByIds

}

func (usergroup *UserGroup) GetUserGroupByName() (UserGroupById, bool) {
	db := GetDbInstance()
	db.First(usergroup, "group_name = ?", usergroup.GroupName)
	userGroupbyId := UserGroupById{GroupName: usergroup.GroupName, GroupID: usergroup.ID}
	if db.RecordNotFound() == true {
		return userGroupbyId, false
	}
	return userGroupbyId, true
}

func (usergroup *UserGroup) GetUserGroupById() (UserGroupById, bool) {
	db := GetDbInstance()
	db.First(usergroup, usergroup.ID)
	if db.RecordNotFound() == true {
		return UserGroupById{GroupID: usergroup.ID, GroupName: usergroup.GroupName}, false
	}
	return UserGroupById{GroupID: usergroup.ID, GroupName: usergroup.GroupName}, true
}
