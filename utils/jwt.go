package utils

//jwt
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

import (
	"errors"
	"time"
)

var jwtKey = []byte(viper.GetString("jwt.JwtKey"))

// 一些常量
var (
	TokenExpired error = errors.New("token is expired")
)

type Claims struct {
	Mobile   string `json:"mobile"`
	PassWord string `json:"pass_word"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(mobile, password string) (string, error) {
	claims := &Claims{
		Mobile:   mobile,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 过期时间
			Issuer:    "storm",                               // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token格式不正确")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token已过期")
			} else {
				return nil, fmt.Errorf("token无效")
			}
		} else {
			return nil, err
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token无效")
	}
}
