package module

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Email    string `gorm:"size:50"`
	Password string `gorm:"size:255"`
	Intro    string `gorm:"default:' '"`
	CreateAt time.Time
	UpdateAt time.Time
}
