package initial

import "github.com/wisaitas/rbac-golang/pkg"

type util struct {
	RedisUtil       pkg.RedisUtil
	JWTUtil         pkg.JWTUtil
	TransactionUtil pkg.TransactionUtil
	ValidatorUtil   pkg.ValidatorUtil
	BcryptUtil      pkg.BcryptUtil
}

func initializeUtil(config *config) *util {
	return &util{
		RedisUtil:       pkg.NewRedisUtil(config.Redis),
		JWTUtil:         pkg.NewJWTUtil(),
		TransactionUtil: pkg.NewTransactionUtil(config.DB),
		ValidatorUtil:   pkg.NewValidatorUtil(),
		BcryptUtil:      pkg.NewBcrypt(),
	}
}
