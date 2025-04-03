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

func initializeValidates(utils *Utils) *Validates {
	return &Validates{
		AddressValidate:        *validates.NewAddressValidate(utils.ValidatorUtil),
		AuthValidate:           *validates.NewAuthValidate(utils.ValidatorUtil),
		DistrictValidate:       *validates.NewDistrictValidate(utils.ValidatorUtil),
		PermissionValidate:     *validates.NewPermissionValidate(utils.ValidatorUtil),
		ProvinceValidate:       *validates.NewProvinceValidate(utils.ValidatorUtil),
		RoleValidate:           *validates.NewRoleValidate(utils.ValidatorUtil),
		RolePermissionValidate: *validates.NewRolePermissionValidate(utils.ValidatorUtil),
		SubDistrictValidate:    *validates.NewSubDistrictValidate(utils.ValidatorUtil),
		UserValidate:           *validates.NewUserValidate(utils.ValidatorUtil),
		UserHistoryValidate:    *validates.NewUserHistoryValidate(utils.ValidatorUtil),
		UserRoleValidate:       *validates.NewUserRoleValidate(utils.ValidatorUtil),
	}
}
