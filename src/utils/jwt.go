package utils

import (
	"errors"
	"time"

	"github.com/fastgox/fastgox-api-starter/src/core/config"
	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明结构
type Claims struct {
	UserID int64  `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

// GenerateJWT 生成JWT令牌
func GenerateJWT(userID int64, phone string) (string, *time.Time, error) {
	if config.GlobalConfig == nil {
		return "", nil, errors.New("配置未初始化")
	}

	// 设置过期时间（24小时）
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建声明
	claims := &Claims{
		UserID: userID,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.GlobalConfig.App.Name,
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名令牌
	tokenString, err := token.SignedString([]byte(config.GlobalConfig.JWT.SecretKey))
	if err != nil {
		return "", nil, err
	}

	return tokenString, &expirationTime, nil
}

// ParseJWT 解析JWT令牌
func ParseJWT(tokenString string) (*Claims, error) {
	if config.GlobalConfig == nil {
		return nil, errors.New("配置未初始化")
	}

	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return []byte(config.GlobalConfig.JWT.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌并提取声明
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌")
}

// ValidateJWT 验证JWT令牌有效性
func ValidateJWT(tokenString string) (*Claims, error) {
	claims, err := ParseJWT(tokenString)
	if err != nil {
		return nil, err
	}

	// 检查令牌是否过期
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("令牌已过期")
	}

	return claims, nil
}
