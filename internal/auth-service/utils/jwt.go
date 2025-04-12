package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wisaitas/rbac-golang/internal/auth-service/env"
	"github.com/wisaitas/rbac-golang/pkg"
)

func GenerateJWTToken(data map[string]interface{}, exp int64) (string, error) {
	claim := jwt.MapClaims(data)
	claim["exp"] = exp
	claim["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(env.ENV.JWT_SECRET))
	if err != nil {
		return "", pkg.Error(err)
	}

	return tokenString, nil
}

func GenerateRedisToken(data map[string]interface{}, exp int64) (string, error) {
	claim := jwt.MapClaims(data)
	claim["exp"] = exp
	claim["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(env.ENV.JWT_SECRET))
	if err != nil {
		return "", pkg.Error(err)
	}

	return tokenString, nil
}
