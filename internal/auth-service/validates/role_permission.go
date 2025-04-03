package validates

import "github.com/wisaitas/rbac-golang/pkg"

type RolePermissionValidate interface {
}

type rolePermissionValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewRolePermissionValidate(validatorUtil pkg.ValidatorUtil) RolePermissionValidate {
	return &rolePermissionValidate{
		validatorUtil: validatorUtil,
	}
}
