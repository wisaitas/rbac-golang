package validates

import "github.com/wisaitas/rbac-golang/pkg"

type RolePermissionValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewRolePermissionValidate(validatorUtil pkg.ValidatorUtil) *RolePermissionValidate {
	return &RolePermissionValidate{
		validatorUtil: validatorUtil,
	}
}
