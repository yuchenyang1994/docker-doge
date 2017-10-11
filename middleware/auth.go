package middleware

import (
	"sync"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	enforcer *casbin.Enforcer
	once     sync.Once
)

const (
	ROLE_SUPPER = "SUPER"
	ROLE_ADMIN  = "ADMIN"
	ROLE_LEADER = "LEADER"
	ROLE_USER   = "USER"
)

// get Auth 单例模式返回权限校验器
func GetAuthzInstance() *casbin.Enforcer {
	once.Do(
		func() {
			adpter := gormadapter.NewAdapter("sqlite3", "data.db", true)
			enforcer = casbin.NewEnforcer("./configs/authz_model.conf", adpter)
		})
	return enforcer
}
