package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Userid   uint   `json:"userid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var MySigningKey = []byte("xiangxun666")

func GenerateJWT(userid uint, username string, role string) (string, error) {
	claims := MyCustomClaims{
		Userid:   userid,   // 用户id
		Username: username, // 用户名
		Role:     role,     // 用户角色（默认为admin）
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "myblog-server",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err

	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (any, error) {
		return MySigningKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("未知的claims type, cannot proceed")
	}
	return claims, nil
}
