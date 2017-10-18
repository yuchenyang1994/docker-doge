package main

import (
	"docker-doge/db"
	"docker-doge/middleware"
	"docker-doge/middleware/policy"
	"strconv"
)

// 迁移数据库
func migrate() {
	d := db.GetDbInstance()
	db.MigrationDB(d)
}

// 同步权限策略
func migratePolicy() {
	ug := db.UserGroup{}
	userGroups := ug.GetUserGroups()
	for _, usergroup := range userGroups {
		if usergroup.GroupName != "ROOT" {
			policy.AddPolicyForUserGroups(usergroup.GroupName)
		}
	}
}

// 创建超级用户
func createRoot() {
	d := db.GetDbInstance()
	ug := db.UserGroup{GroupName: "ROOT"}
	d.Create(&ug)
	user := &db.User{UserGroupID: ug.ID, Email: "admin@gmail.com", Password: "admin"}
	user.Insert(d)
	e := middleware.GetAuthzInstance()
	strUserId := strconv.Itoa(int(user.ID))
	e.AddRoleForUser(strUserId, "SUPER")
}
