package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义一个密钥，用于签名和验证 JWT
var mySigningKey = []byte("secret")

type MyCustomClaims struct {
	UserId int32 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId int32) string {
	// 设置 Claims
	claims := MyCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 过期时间
			Issuer:    "hahaha3w",                                // 签发者
		},
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的 JWT
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return ""
	}

	return tokenString
}
