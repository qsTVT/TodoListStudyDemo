package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTsecret = []byte("todo_list_jwt_secret")

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{Id: id, Username: username, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "todo_list",
	}}
	// 生成令牌：使用HS256算法签名
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
