package db

import (
	"crypto/sha1"

	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserGroupID uint
	Email       string `gorm:"size:50"`
	Password    string `gorm:"size:255"`
	Intro       string `gorm:"default:' '"`
}

func (user *User) HashPassword() {
	hash := sha1.New()
	hash.Write([]byte(user.Password))
	hashPassword := fmt.Sprintf("SHA1_%x", hash.Sum(nil))
	user.Password = hashPassword
}

func (user *User) Insert(db *gorm.DB) error {
	defer db.Close()
	user.HashPassword()
	db.Create(user)
	return db.Error
}

func (user *User) GetUserByPassword(db *gorm.DB) (*User, bool) {
	defer db.Close()
	user.HashPassword()
	u := db.Where("email=? AND password=?", user.Email, user.Password).First(user)
	if u.RecordNotFound() == true {
		return user, false
	} else {
		return user, true
	}
}
