package middleware

import (
	"sync"
	"time"

	"strconv"

	"docker-doge/db"

	"docker-doge/configs"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

// GetAuthzInstance 权限校验器
func GetAuthzInstance() *casbin.Enforcer {
	once.Do(func() {
		conf := configs.Conf()
		adpter := gormadapter.NewAdapter(conf.DATABASE_BACKEND, conf.DATABASE_URI, true)
		enforcer = casbin.NewEnforcer("./configs/authz_model.conf", adpter)
	})
	return enforcer
}

// NewJwtMiddleWare a GinJwtMiddleWare
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

// JwtAuthenticatorHandler get UserId
func JwtAuthenticatorHandler(username string, password string, c *gin.Context) (string, bool) {
	d := db.GetDbInstance()
	user := &db.User{Email: username, Password: password}
	user, has := user.GetUserByPassword(d)
	userID := strconv.Itoa(int(user.ID))
	if has {
		return userID, true
	}
	return userID, false
}

// JwtAuthorizatorHandler checkPermission for User
func JwtAuthorizatorHandler(userID string, c *gin.Context) bool {
	d := db.GetDbInstance()
	uUserID, _ := strconv.ParseUint(userID, 0, 64)
	user := db.User{}
	d.First(&user, uUserID)
	e := GetAuthzInstance()
	userRoles := e.GetRolesForUser(userID)
	return checkPermission(c, e, d, &user, userRoles)
}

// JwtUnauthorized on error
func JwtUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func checkPermission(c *gin.Context, enforcer *casbin.Enforcer, d *gorm.DB, user *db.User, roles []string) bool {
	r := c.Request
	method := r.Method
	path := r.URL.Path
	usergroup := db.UserGroup{}
	d.Find(&usergroup, user.UserGroupID)
	for _, role := range roles {
		return enforcer.Enforce(usergroup.GroupName, role, path, method)
	}
	return false
}
