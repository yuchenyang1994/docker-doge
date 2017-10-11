package configs

import (
	"docker-doge/handler"
	"time"

	jwt "github.com/appleboy/gin-jwt"
)

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
