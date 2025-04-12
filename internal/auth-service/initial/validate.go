package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/validates"
)

type validate struct {
	AddressValidate        validates.AddressValidate
	AuthValidate           validates.AuthValidate
	DistrictValidate       validates.DistrictValidate
	PermissionValidate     validates.PermissionValidate
	ProvinceValidate       validates.ProvinceValidate
	RoleValidate           validates.RoleValidate
	RolePermissionValidate validates.RolePermissionValidate
	SubDistrictValidate    validates.SubDistrictValidate
	UserValidate           validates.UserValidate
	UserHistoryValidate    validates.UserHistoryValidate
	UserRoleValidate       validates.UserRoleValidate
}

func initializeValidate(util *util) *validate {
	return &validate{
		AddressValidate:        validates.NewAddressValidate(util.ValidatorUtil),
		AuthValidate:           validates.NewAuthValidate(util.ValidatorUtil),
		DistrictValidate:       validates.NewDistrictValidate(util.ValidatorUtil),
		PermissionValidate:     validates.NewPermissionValidate(util.ValidatorUtil),
		ProvinceValidate:       validates.NewProvinceValidate(util.ValidatorUtil),
		RoleValidate:           validates.NewRoleValidate(util.ValidatorUtil),
		RolePermissionValidate: validates.NewRolePermissionValidate(util.ValidatorUtil),
		SubDistrictValidate:    validates.NewSubDistrictValidate(util.ValidatorUtil),
		UserValidate:           validates.NewUserValidate(util.ValidatorUtil),
		UserHistoryValidate:    validates.NewUserHistoryValidate(util.ValidatorUtil),
		UserRoleValidate:       validates.NewUserRoleValidate(util.ValidatorUtil),
	}
}
