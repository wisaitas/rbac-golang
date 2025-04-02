package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/validates"
)

type Validates struct {
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

func initializeValidates() *Validates {
	return &Validates{
		AddressValidate:        *validates.NewAddressValidate(),
		AuthValidate:           *validates.NewAuthValidate(),
		DistrictValidate:       *validates.NewDistrictValidate(),
		PermissionValidate:     validates.NewPermissionValidate(),
		ProvinceValidate:       *validates.NewProvinceValidate(),
		RoleValidate:           *validates.NewRoleValidate(),
		RolePermissionValidate: *validates.NewRolePermissionValidate(),
		SubDistrictValidate:    *validates.NewSubDistrictValidate(),
		UserValidate:           *validates.NewUserValidate(),
		UserHistoryValidate:    *validates.NewUserHistoryValidate(),
		UserRoleValidate:       *validates.NewUserRoleValidate(),
	}
}
