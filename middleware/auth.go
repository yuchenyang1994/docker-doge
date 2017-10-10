package middleware

import (
	"docker-doge/handler"
	"github.com/appleboy/gin-jwt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
	"time"
)

var (
	enforcer *casbin.Enforcer
	once     sync.Once
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

func NewJwtMiddleWare() *jwt.GinJWTMiddleware {
	// 此处记得增加配置文件
	j := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: handler.JwtAuthenticatorHandler,
		Authorizator:  handler.JwtAuthorizatorHandler,
		Unauthorized:  handler.JwtUnauthorized,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	return j
}
