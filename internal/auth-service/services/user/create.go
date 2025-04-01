package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
	"golang.org/x/crypto/bcrypt"
)

type Create interface {
	CreateUser(req requests.CreateUserRequest) (resp responses.CreateUserResponse, statusCode int, err error)
	AssignRole(req requests.AssignRoleRequest) (resp responses.UsersResponse, statusCode int, err error)
}

type create struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewCreate(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Create {
	return &create{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *create) CreateUser(req requests.CreateUserRequest) (resp responses.CreateUserResponse, statusCode int, err error) {
	user := req.ReqToModel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return resp, http.StatusInternalServerError, err
	}

	user.Password = string(hashedPassword)

	if err = r.userRepository.Create(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return resp, http.StatusBadRequest, errors.New("username already exists")
		}

		return resp, http.StatusInternalServerError, err
	}

	return resp.ModelToResponse(user), http.StatusCreated, nil
}

func (r *create) AssignRole(req requests.AssignRoleRequest) (resp responses.UsersResponse, statusCode int, err error) {
	// TODO: Implement
	return resp, http.StatusOK, nil
}
