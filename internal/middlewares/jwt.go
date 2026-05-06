package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xmujin/myblog-backend/pkg/auth"
)

// jwt认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"获取文章失败": "没有token"})
			c.Abort() // 直接拦截
			return
		}

		// 删除 "Bearer "前缀
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &auth.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return auth.MySigningKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"获取文章失败": "非法的token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*auth.MyCustomClaims); ok {
			c.Set("userid", claims.Userid)
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)
			c.Next()
		}
	}
}
