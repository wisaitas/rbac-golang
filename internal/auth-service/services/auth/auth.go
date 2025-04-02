package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/internal/auth-service/utils"
	"github.com/wisaitas/rbac-golang/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(req requests.LoginRequest) (resp responses.LoginResponse, statusCode int, err error)
	Register(req requests.RegisterRequest) (resp responses.RegisterResponse, statusCode int, err error)
	Logout(tokenContext models.TokenContext) (statusCode int, err error)
	RefreshToken(userContext models.UserContext) (resp responses.LoginResponse, statusCode int, err error)
}

type authService struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	roleRepository        repositories.RoleRepository
	redis                 pkg.RedisUtil
	jwtUtil               pkg.JWTUtil
}

func NewAuthService(
	userRepository repositories.UserRepository,
	userHistoryRepository repositories.UserHistoryRepository,
	roleRepository repositories.RoleRepository,
	redis pkg.RedisUtil,
	jwtUtil pkg.JWTUtil,
) AuthService {
	return &authService{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		roleRepository:        roleRepository,
		redis:                 redis,
		jwtUtil:               jwtUtil,
	}
}

func (r *authService) Login(req requests.LoginRequest) (resp responses.LoginResponse, statusCode int, err error) {
	user := models.User{}

	relations := []pkg.Relation{
		{
			Query: "Roles",
		},
	}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("username = ?", req.Username), &relations); err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, http.StatusNotFound, pkg.Error(err)
		}

		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return resp, http.StatusUnauthorized, pkg.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(20 * time.Minute)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	tokenData := map[string]interface{}{
		"user_id": user.ID,
	}

	accessToken, err := utils.GenerateJWTToken(tokenData, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	refreshToken, err := utils.GenerateJWTToken(tokenData, refreshTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	sort.Slice(user.Roles, func(i, j int) bool {
		return user.Roles[i].Priority < user.Roles[j].Priority
	})

	role := user.Roles[0]

	tokenRedisData := models.UserContext{
		UserID:       user.ID,
		Username:     user.Username,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		RoleID:       role.ID,
		RoleName:     role.Name,
	}

	tokenRedisDataJson, err := json.Marshal(tokenRedisData)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), string(tokenRedisDataJson), accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), string(tokenRedisDataJson), refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(accessToken, refreshToken), statusCode, pkg.Error(err)
}

func (r *authService) Register(req requests.RegisterRequest) (resp responses.RegisterResponse, statusCode int, err error) {
	user := req.ReqToModel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	user.Password = string(hashedPassword)

	role := models.Role{}
	if err := r.roleRepository.GetBy(&role, pkg.NewCondition("name = ?", "User"), nil); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	user.Roles = append(user.Roles, role)

	if err = r.userRepository.Create(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return resp, http.StatusBadRequest, pkg.Error(errors.New("username already exists"))
		}

		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ModelToResponse(user), http.StatusCreated, pkg.Error(err)
}

func (r *authService) Logout(tokenContext models.TokenContext) (statusCode int, err error) {
	if err := r.redis.Del(context.Background(), fmt.Sprintf("access_token:%s", tokenContext.UserID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Del(context.Background(), fmt.Sprintf("refresh_token:%s", tokenContext.UserID)); err != nil {
		return http.StatusInternalServerError, pkg.Error(err)
	}

	return http.StatusOK, nil
}

func (r *authService) RefreshToken(userContext models.UserContext) (resp responses.LoginResponse, statusCode int, err error) {
	user := models.User{}

	relations := []pkg.Relation{
		{
			Query: "Roles",
		},
	}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("username = ?", userContext.Username), &relations); err != nil {
		return resp, http.StatusNotFound, pkg.Error(err)
	}

	timeNow := time.Now()
	accessTokenExp := timeNow.Add(20 * time.Minute)
	refreshTokenExp := timeNow.Add(time.Hour * 24)

	tokenData := map[string]interface{}{
		"id": user.ID,
	}

	accessToken, err := utils.GenerateJWTToken(tokenData, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	refreshToken, err := utils.GenerateJWTToken(tokenData, refreshTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	tokenRedisData := map[string]interface{}{
		"id":            user.ID,
		"username":      user.Username,
		"email":         user.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	accessTokenRedis, err := utils.GenerateRedisToken(tokenRedisData, accessTokenExp.Unix())
	if err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("access_token:%s", user.ID), accessTokenRedis, accessTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redis.Set(context.Background(), fmt.Sprintf("refresh_token:%s", user.ID), refreshToken, refreshTokenExp.Sub(timeNow)); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ToResponse(accessToken, refreshToken), http.StatusOK, nil
}
