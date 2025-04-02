package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Get interface {
	GetUsers(query pkg.PaginationQuery) (resp []responses.UsersResponse, statusCode int, err error)
	GetUserProfile(userContext models.UserContext) (resp responses.UsersResponse, statusCode int, err error)
}

type get struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewGet(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Get {
	return &get{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *get) GetUsers(query pkg.PaginationQuery) (resp []responses.UsersResponse, statusCode int, err error) {
	users := []models.User{}

	cacheKey := fmt.Sprintf("get_users:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.userRepository.GetAll(&users, &query, nil, nil); err != nil {
		return []responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, user := range users {
		respUser := responses.UsersResponse{}
		resp = append(resp, respUser.ModelToResponse(user))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}

func (r *get) GetUserProfile(userContext models.UserContext) (resp responses.UsersResponse, statusCode int, err error) {
	user := models.User{}

	cacheKey := fmt.Sprintf("get_user_profile:%v", userContext.UserID)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	relations := []pkg.Relation{
		{
			Query: "Addresses",
		},
		{
			Query: "Roles",
		},
	}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("id = ?", userContext.UserID), &relations); err != nil {
		return responses.UsersResponse{}, http.StatusNotFound, pkg.Error(err)
	}

	resp.ModelToResponse(user)

	respJson, err := json.Marshal(resp)
	if err != nil {
		return responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return responses.UsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}
