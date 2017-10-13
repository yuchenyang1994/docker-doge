package main

import "docker-doge/db"

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
