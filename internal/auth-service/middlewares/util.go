package middlewares

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/rbac-golang/internal/auth-service/configs"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

func authToken(c *fiber.Ctx, redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return pkg.Error(errors.New("invalid token type"))
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var userContext models.UserContext
	_, err := jwt.ParseWithClaims(token, &userContext, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.ENV.JWT_SECRET), nil
	})
	if err != nil {
		return pkg.Error(err)
	}

	_, err = redisUtil.Get(context.Background(), fmt.Sprintf("access_token:%s", uuid.MustParse(userContext.ID)))
	if err != nil {
		if err == redis.Nil {
			return pkg.Error(errors.New("session not found"))
		}

		return pkg.Error(err)
	}

	c.Locals("userContext", userContext)
	return nil
}

func authRefreshToken(c *fiber.Ctx, redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return pkg.Error(errors.New("invalid token type"))
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var userContext models.UserContext
	_, err := jwt.ParseWithClaims(token, &userContext, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.ENV.JWT_SECRET), nil
	})
	if err != nil {
		return pkg.Error(err)
	}

	_, err = redisUtil.Get(context.Background(), fmt.Sprintf("refresh_token:%s", uuid.MustParse(userContext.ID)))
	if err != nil {
		if err == redis.Nil {
			return pkg.Error(errors.New("session not found"))
		}

		return pkg.Error(err)
	}

	c.Locals("userContext", userContext)
	return nil
}
