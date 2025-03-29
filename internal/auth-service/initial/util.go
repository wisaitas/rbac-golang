package initial

import "github.com/wisaitas/rbac-golang/pkg"

type Utils struct {
	RedisUtil       pkg.RedisUtil
	JWTUtil         pkg.JWTUtil
	TransactionUtil pkg.TransactionUtil
	ValidatorUtil   pkg.ValidatorUtil
	BcryptUtil      pkg.BcryptUtil
}

func initializeUtils(configs *Configs) *Utils {
	return &Utils{
		RedisUtil:       pkg.NewRedisUtil(configs.Redis),
		JWTUtil:         pkg.NewJWTUtil(),
		TransactionUtil: pkg.NewTransactionUtil(configs.DB),
		ValidatorUtil:   pkg.NewValidatorUtil(),
		BcryptUtil:      pkg.NewBcrypt(),
	}
}
