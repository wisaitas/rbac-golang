package validates

import "github.com/wisaitas/rbac-golang/pkg"

type UserHistoryValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewUserHistoryValidate(validatorUtil pkg.ValidatorUtil) *UserHistoryValidate {
	return &UserHistoryValidate{
		validatorUtil: validatorUtil,
	}
}
