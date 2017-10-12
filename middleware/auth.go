package middleware

import (
	"net/http"
	"sync"

	"strconv"

	"docker-doge/db"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
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
	adpter := gormadapter.NewAdapter("sqlite3", "data.db", true)
	enforcer = casbin.NewEnforcer("./configs/authz_model.conf", adpter)
	return enforcer
}

// NewJwtAuthorizer returns the authorizer, uses a Casbin enforcer as gin-JWT
func NewJwtAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		a := &BasicAuthorizer{enforcer: e}

		if !a.CheckPermission(c.Request, c) {
			a.RequirePermission(c.Writer)
		}
	}
}

// JwtAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the username from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserRole(c *gin.Context) []string {
	d := db.GetDbInstance()
	claims := jwt.ExtractClaims(c)
	strUserId := claims["id"].(string)
	user := &db.User{}
	userid, _ := strconv.ParseUint(strUserId, 0, 64)
	d.First(user, userid)
	roles := a.enforcer.GetRolesForUser(user.Email)
	return roles
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request, c *gin.Context) bool {
	roles := a.GetUserRole(c)
	method := r.Method
	path := r.URL.Path
	for _, role := range roles {
		if a.enforcer.Enforce(role, path, method) {
			return true
		}
	}
	return false
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(w http.ResponseWriter) {
	w.WriteHeader(403)
	w.Write([]byte("403 Forbidden\n"))
}
