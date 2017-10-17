package main

import "docker-doge/db"
import "docker-doge/handler"

// 迁移数据库
func migrate() {
	d := db.GetDbInstance()
	db.MigrationDB(d)
}

// 生成测试usergroup
func createUserGroup() {
	d := db.GetDbInstance()
	db.CreateUserGroup(d)
}

func migratePolicy() {
	ug := db.UserGroup{}
	userGroups := ug.GetUserGroups()
	for _, usergroup := range userGroups {
		handler.AddPolicyForUserGroups(usergroup.GroupName)
	}
}
