package jwtutil

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	v1 "github.com/ividernvi/algohub/model/v1"
)

// SecretKey 用于签名 JWT 的密钥
var (
	REALM          = os.Getenv("ALGOHUB_JWT_REALM")
	SecretKey      = []byte(os.Getenv("ALGOHUB_JWT_SECRET"))
	ExpireDuration = time.Duration(2 * time.Hour)
)

// CreateJWT 创建一个 JWT
func CreateJWT(user *v1.User) (string, error) {
	// 创建一个声明
	claims := jwt.MapClaims{
		"username":      user.UserName,
		"user_instance": user.InstanceID,
		"user_status":   user.Status,
		"exp":           time.Now().Add(ExpireDuration).Unix(),
		"realm":         REALM,
	}

	// 创建一个带有声明的 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名 token
	return token.SignedString(SecretKey)
}

// ValidateJWT 验证并解析 JWT
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// 解析并验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保使用的是正确的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 提取声明
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
