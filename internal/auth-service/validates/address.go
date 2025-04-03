package validates

import "github.com/wisaitas/rbac-golang/pkg"

type AddressValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewAddressValidate(validatorUtil pkg.ValidatorUtil) *AddressValidate {
	return &AddressValidate{
		validatorUtil: validatorUtil,
	}
}
