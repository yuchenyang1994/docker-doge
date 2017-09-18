package db

import (
	"github.com/jinzhu/gorm"
)

// 用户组
type UserGroup struct {
	gorm.Model
	ID        int
	GroupName string `gorm:"size:60"`
	Users     []User
}
