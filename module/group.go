package module

import (
	"github.com/jinzhu/gorm"
)

type UserGroup struct {
	gorm.Model
	ID        int
	GroupName string `gorm:"size:60"`
}
