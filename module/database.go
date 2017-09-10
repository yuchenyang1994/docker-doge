package module

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
		db, _ = gorm.Open("sqlite3", "./data.db")
	})
	return db
}
