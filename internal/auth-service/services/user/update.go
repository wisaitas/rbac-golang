package user

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/constants"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/params"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
	"gorm.io/gorm"
)

type Update interface {
	UpdateUser(param params.UserParams, req requests.UpdateUserRequest) (resp responses.UpdateUserResponse, statusCode int, err error)
}

type update struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	redisUtil             pkg.RedisUtil
	transactionUtil       pkg.TransactionUtil
}

func NewUpdate(
	userRepository repositories.UserRepository,
	userHistoryRepository repositories.UserHistoryRepository,
	redisUtil pkg.RedisUtil,
	transactionUtil pkg.TransactionUtil,
) Update {
	return &update{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redisUtil:             redisUtil,
		transactionUtil:       transactionUtil,
	}
}

func (r *update) UpdateUser(param params.UserParams, request requests.UpdateUserRequest) (resp responses.UpdateUserResponse, statusCode int, err error) {
	user := models.User{}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("id = ?", param.ID), "Addresses"); err != nil {
		return resp, http.StatusNotFound, pkg.Error(err)
	}

	if err := r.transactionUtil.ExecuteInTransaction(func(tx *gorm.DB) error {
		txUserRepository := r.userRepository.WithTx(tx)
		txUserHistoryRepository := r.userHistoryRepository.WithTx(tx)

		userBeforeUpdate := models.UserHistory{
			Action:       constants.Action.Update,
			UserID:       user.ID,
			OldFirstName: user.FirstName,
			OldLastName:  user.LastName,
			OldBirthDate: user.BirthDate,
			OldPassword:  user.Password,
			OldVersion:   user.Version,
		}

		if err := txUserHistoryRepository.Create(&userBeforeUpdate); err != nil {
			return pkg.Error(err)
		}

		if request.FirstName != nil {
			user.FirstName = *request.FirstName
		}

		if request.LastName != nil {
			user.LastName = *request.LastName
		}

		if request.BirthDate != nil {
			user.BirthDate = *request.BirthDate
		}

		if err := txUserRepository.Update(&user); err != nil {
			return pkg.Error(err)
		}

		return nil
	}); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ModelToResponse(user), http.StatusOK, nil
}
