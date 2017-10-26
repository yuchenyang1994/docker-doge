package db

import (
	"log"

	"docker-doge/configs"

	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db   *gorm.DB
	once sync.Once
	err  error
)

func GetDbInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("db").(*gorm.DB)
}

func CreateDb() *gorm.DB {
	once.Do(func() {
		conf := configs.Conf()
		db, err = gorm.Open(conf.DATABASE_BACKEND, conf.DATABASE_URI)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		if err != nil {
			log.Fatal("db error")
		}
	})
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
	u := UserGroup{GroupName: "ROOT"}
	db.NewRecord(u)
	db.Create(&u)
}

func DataBase() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := CreateDb()
		c.Set("db", db)
		c.Next()
		db.Close()
	}

}
