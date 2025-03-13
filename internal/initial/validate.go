package initial

import (
	"github.com/wisaitas/standard-golang/internal/validates"
)

type Validates struct {
	AddressValidate         validates.AddressValidate
	AuthValidate            validates.AuthValidate
	CategoryValidate        validates.CategoryValidate
	DistrictValidate        validates.DistrictValidate
	OrderValidate           validates.OrderValidate
	OrderItemValidate       validates.OrderItemValidate
	PermissionValidate      validates.PermissionValidate
	ProductValidate         validates.ProductValidate
	ProductImageValidate    validates.ProductImageValidate
	ProvinceValidate        validates.ProvinceValidate
	RoleValidate            validates.RoleValidate
	RolePermissionValidate  validates.RolePermissionValidate
	ShippingAddressValidate validates.ShippingAddressValidate
	SubDistrictValidate     validates.SubDistrictValidate
	UserValidate            validates.UserValidate
	UserHistoryValidate     validates.UserHistoryValidate
	UserRoleValidate        validates.UserRoleValidate
	WishlistValidate        validates.WishlistValidate
}

func initializeValidates() *Validates {
	return &Validates{
		AddressValidate:         *validates.NewAddressValidate(),
		AuthValidate:            *validates.NewAuthValidate(),
		CategoryValidate:        *validates.NewCategoryValidate(),
		DistrictValidate:        *validates.NewDistrictValidate(),
		OrderValidate:           *validates.NewOrderValidate(),
		OrderItemValidate:       *validates.NewOrderItemValidate(),
		PermissionValidate:      *validates.NewPermissionValidate(),
		ProductValidate:         *validates.NewProductValidate(),
		ProductImageValidate:    *validates.NewProductImageValidate(),
		ProvinceValidate:        *validates.NewProvinceValidate(),
		RoleValidate:            *validates.NewRoleValidate(),
		RolePermissionValidate:  *validates.NewRolePermissionValidate(),
		ShippingAddressValidate: *validates.NewShippingAddressValidate(),
		SubDistrictValidate:     *validates.NewSubDistrictValidate(),
		UserValidate:            *validates.NewUserValidate(),
		UserHistoryValidate:     *validates.NewUserHistoryValidate(),
		UserRoleValidate:        *validates.NewUserRoleValidate(),
		WishlistValidate:        *validates.NewWishlistValidate(),
	}
}
