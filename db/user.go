package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserGroupID int
	RoleID      int
	Email       string `gorm:"size:50"`
	Password    string `gorm:"size:255"`
	Intro       string `gorm:"default:' '"`
}
