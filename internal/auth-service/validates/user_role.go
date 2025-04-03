package validates

import "github.com/wisaitas/rbac-golang/pkg"

type UserRoleValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewUserRoleValidate(validatorUtil pkg.ValidatorUtil) *UserRoleValidate {
	return &UserRoleValidate{
		validatorUtil: validatorUtil,
	}
}
