package services

import (
	"docker-doge/db"

	"strconv"

	"github.com/casbin/casbin"
)

type (
	JsonUserInfos struct {
		UserGroupName string     `json:"groupName"`
		Users         []UserInfo `json:"users"`
	}
	UserInfo struct {
		Roles []string `json:"roles"`
		Email string   `json:"email"`
		ID    int      `json:"id"`
	}

	UserInfoService struct {
		Enforcer  *casbin.Enforcer
		GroupName string
	}
)

func NewUserInfoService(e *casbin.Enforcer, groupName string) *UserInfoService {
	service := UserInfoService{e, groupName}
	return &service
}

// GetUserInfos ...
func (service *UserInfoService) GetUserInfos() []JsonUserInfos {
	d := db.GetDbInstance()
	if service.GroupName == "Super" {
		usergroups := []db.UserGroup{}
		d.Find(&usergroups)
		jsonUserInfos := []JsonUserInfos{}
		for _, usergroup := range usergroups {
			users := []db.User{}
			d.Model(&usergroup).Related(&users)
			userinfos := createUserInfo(users, service.Enforcer)
			jsonUserInfo := JsonUserInfos{UserGroupName: usergroup.GroupName, Users: userinfos}
			jsonUserInfos = append(jsonUserInfos, jsonUserInfo)
		}
		return jsonUserInfos
	}
	usergroups := []db.UserGroup{}
	d.Find(&usergroups, "group_name = ?", service.GroupName)
	jsonUserInfos := []JsonUserInfos{}
	for _, usergroup := range usergroups {
		users := []db.User{}
		d.Model(&usergroup).Related(&users)
		userinfos := createUserInfo(users, service.Enforcer)
		jsonUserInfo := JsonUserInfos{UserGroupName: usergroup.GroupName, Users: userinfos}
		jsonUserInfos = append(jsonUserInfos, jsonUserInfo)
	}
	return jsonUserInfos

}

// createUserInfo ...
func createUserInfo(users []db.User, e *casbin.Enforcer) []UserInfo {
	userinfoList := []UserInfo{}
	for _, user := range users {
		userId := strconv.Itoa(int(user.ID))
		roles := e.GetRolesForUser(userId)
		userinfo := UserInfo{roles, user.Email, int(user.ID)}
		userinfoList = append(userinfoList, userinfo)
	}
	return userinfoList
}
