package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDbInstance() *gorm.DB {
	once.Do(func() {
		db, _ = gorm.Open("sqlite3", "data.db")
	})
	return db
}

// 迁移数据库模型函数
func MigrationDB(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&UserGroup{},
		&Role{})
}

// 创建基本的角色数据
// CreateRole ...
func CreateRole(db *gorm.DB) {
	role := Role{}
	roleList := []string{"super", "admin", "leader", "user"}
	for _, value := range roleList {
		db.Where("name = ?", value).Find(&role)
		if role.Name == "" {
			db.NewRecord(Role{Name: value})
		}
	}
}
