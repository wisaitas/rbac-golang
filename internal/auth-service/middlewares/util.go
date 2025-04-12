package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/rbac-golang/internal/auth-service/env"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
)

func authToken(c *fiber.Ctx, redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return pkg.Error(errors.New("invalid token type"))
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var tokenContext models.TokenContext
	if err := jwtUtil.ValidateToken(token, &tokenContext, env.ENV.JWT_SECRET); err != nil {
		return pkg.Error(err)
	}

	userContextJson, err := redisUtil.Get(context.Background(), fmt.Sprintf("access_token:%s", tokenContext.UserID))
	if err != nil {
		if err == redis.Nil {
			return pkg.Error(errors.New("session not found"))
		}

		return pkg.Error(err)
	}

	userContext := models.UserContext{}
	if err := json.Unmarshal([]byte(userContextJson), &userContext); err != nil {
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

	var tokenContext models.TokenContext
	if err := jwtUtil.ValidateToken(token, &tokenContext, env.ENV.JWT_SECRET); err != nil {
		return pkg.Error(err)
	}

	userContextJson, err := redisUtil.Get(context.Background(), fmt.Sprintf("refresh_token:%s", tokenContext.UserID))
	if err != nil {
		if err == redis.Nil {
			return pkg.Error(errors.New("session not found"))
		}

		return pkg.Error(err)
	}

	userContext := models.UserContext{}
	if err := json.Unmarshal([]byte(userContextJson), &userContext); err != nil {
		return pkg.Error(err)
	}

	c.Locals("userContext", userContext)
	return nil
}
