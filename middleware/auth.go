package middleware

import (
	"net/http"
	"sync"
	"time"

	"strconv"

	"docker-doge/db"

	"docker-doge/configs"

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

// 权限校验器
func GetAuthzInstance() *casbin.Enforcer {
	conf := configs.Conf()
	adpter := gormadapter.NewAdapter(conf.DATABASE_BACKEND, conf.DATABASE_URI, true)
	enforcer = casbin.NewEnforcer("./configs/authz_model.conf", adpter)
	return enforcer
}

func NewJwtMiddleWare() *jwt.GinJWTMiddleware {
	conf := configs.Conf()
	j := &jwt.GinJWTMiddleware{
		Realm:         conf.REALM,
		Key:           []byte(conf.SCRETKEY),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: JwtAuthenticatorHandler,
		Authorizator:  JwtAuthorizatorHandler,
		Unauthorized:  JwtUnauthorized,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	return j
}

func JwtAuthenticatorHandler(username string, password string, c *gin.Context) (string, bool) {
	d := db.GetDbInstance()
	user := &db.User{Email: username, Password: password}
	user, has := user.GetUserByPassword(d)
	userId := strconv.Itoa(int(user.ID))
	if has {
		return userId, true
	}
	return userId, false
}

func JwtAuthorizatorHandler(userId string, c *gin.Context) bool {
	d := db.GetDbInstance()
	defer d.Close()
	user := &db.User{}
	uuserId, _ := strconv.ParseUint(userId, 0, 64)
	if notFound := d.First(user, uuserId).RecordNotFound(); notFound != true {
		return true
	}
	return false
}

func JwtUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
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
type RoleBaseAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the username from the request.
// Currently, only HTTP basic authentication is supported
func (a *RoleBaseAuthorizer) GetUserRole(c *gin.Context) (*db.User, []string) {
	d := db.GetDbInstance()
	claims := jwt.ExtractClaims(c)
	strUserId := claims["id"].(string)
	user := &db.User{}
	userid, _ := strconv.ParseUint(strUserId, 0, 64)
	d.First(user, userid)
	roles := a.enforcer.GetRolesForUser(user.Email)
	return user, roles
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *RoleBaseAuthorizer) CheckPermission(r *http.Request, c *gin.Context) bool {
	user, roles := a.GetUserRole(c)
	method := r.Method
	path := r.URL.Path
	d := db.GetDbInstance()
	usergroup := db.UserGroup{}
	d.Find(&usergroup, user.UserGroupID)
	for _, role := range roles {
		if a.enforcer.Enforce(groupname, role, path, method) {
			return true
		}
	}
	return false
}

// RequirePermission returns the 403 Forbidden to the client
func (a *RoleBaseAuthorizer) RequirePermission(w http.ResponseWriter) {
	w.WriteHeader(403)
	w.Write([]byte("403 Forbidden\n"))
}
