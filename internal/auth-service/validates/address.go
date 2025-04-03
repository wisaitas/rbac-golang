package validates

import "github.com/wisaitas/rbac-golang/pkg"

type AddressValidate interface {
}

type addressValidate struct {
	validatorUtil pkg.ValidatorUtil
}

func NewAddressValidate(validatorUtil pkg.ValidatorUtil) AddressValidate {
	return &addressValidate{
		validatorUtil: validatorUtil,
	}
}
