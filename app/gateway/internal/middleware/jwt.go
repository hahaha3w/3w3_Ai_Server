package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
)

// 定义一个密钥，用于签名和验证 JWT
var mySigningKey = []byte("secret")

type MyCustomClaims struct {
	UserId int32 `json:"user_id"`
	jwt.StandardClaims
}

func ParseJWT(tokenString string) (*MyCustomClaims, error) {
	// 解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 Token
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			domain.Error(ctx, 403, "Token 为空")
			return
		}

		// 检查 Authorization 头是否以 "Bearer " 开头
		const bearerPrefix = "Bearer "
		if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
			// 去掉 "Bearer " 前缀，获取实际的 Token
			tokenString := authHeader[len(bearerPrefix):]
			claims, err := ParseJWT(tokenString)
			if err != nil {
				domain.Error(ctx, 403, "Token 验证错误")
				return
			}

			ctx.Set("userId", claims.UserId)
			ctx.Next()
		} else {
			domain.Error(ctx, 403, "Token 格式错误")
			return
		}
	}
}
