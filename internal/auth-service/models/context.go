package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenContext struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type UserContext struct {
	UserID       uuid.UUID `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	RoleID       uuid.UUID `json:"role_id"`
	RoleName     string    `json:"role_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
