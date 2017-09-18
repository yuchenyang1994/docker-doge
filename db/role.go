package db

import "github.com/jinzhu/gorm"

// 角色 如管理员，组长，超级管理员之类的
type Role struct {
	gorm.Model
	ID    int
	Name  string `gorm:"size:50"`
	Users []User
}
