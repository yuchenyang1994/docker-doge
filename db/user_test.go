package db

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

// TestUserHashPasswd ...
func Test_UserHashPasswd(t *testing.T) {
	u := &User{Email: "xxx@gamil.com", Password: "123456"}
	hash := sha1.New()
	hash.Write([]byte("123456"))
	hashPassword := fmt.Sprintf("SHA1_%x", hash.Sum(nil))
	u.HashPassword()
	t.Log(hashPassword)
	if u.Password != hashPassword {
		t.Error("password error")
	}
}

func Test_UserInsert(t *testing.T) {
	d := GetTestDB()
	MigrationDB(d)
	tx := d.Begin()
	defer tx.Rollback()
	ug := &UserGroup{GroupName: "test"}
	tx.NewRecord(ug)
	user := &User{Email: "xxx@gmail.com", Password: "123456",
		UserGroupID: ug.ID}
	user.Insert(d)
	if user.UserGroupID != ug.ID {
		t.Error("user group error")
	}
	_, has := user.GetUserByPassword(d)
	if has != true {
		t.Error("user not insted")
	}
}
