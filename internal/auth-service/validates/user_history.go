package validates

import "github.com/wisaitas/rbac-golang/pkg"

type UserHistoryValidate interface {
}

type userHistoryValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewUserHistoryValidate(validatorUtil pkg.ValidatorUtil) UserHistoryValidate {
	return &userHistoryValidate{
		validatorUtil: validatorUtil,
	}
}
