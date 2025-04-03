package validates

import "github.com/wisaitas/rbac-golang/pkg"

type UserRoleValidate interface {
}

type userRoleValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewUserRoleValidate(validatorUtil pkg.ValidatorUtil) UserRoleValidate {
	return &userRoleValidate{
		validatorUtil: validatorUtil,
	}
}
