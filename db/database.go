package db

import (
	"log"

	"docker-doge/configs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetDbInstance() *gorm.DB {
	conf := configs.Conf()
	db, err := gorm.Open(conf.DATABASE_BACKEND, conf.DATABASE_URI)
	if err != nil {
		log.Fatal("db error")
	}
	return db
}

func GetTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		log.Fatal("db error")
	}
	return db
}

// 迁移数据库模型函数
func MigrationDB(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&UserGroup{})
}

func CreateUserGroup(db *gorm.DB) {
	u := UserGroup{GroupName: "test"}
	db.NewRecord(u)
	db.Create(&u)
}
