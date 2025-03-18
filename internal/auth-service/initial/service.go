package initial

import (
	authService "github.com/wisaitas/rbac-golang/internal/auth-service/services/auth"
	districtService "github.com/wisaitas/rbac-golang/internal/auth-service/services/district"
	provinceService "github.com/wisaitas/rbac-golang/internal/auth-service/services/province"
	subDistrictService "github.com/wisaitas/rbac-golang/internal/auth-service/services/sub-district"
	userService "github.com/wisaitas/rbac-golang/internal/auth-service/services/user"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Services struct {
	UserService        userService.UserService
	AuthService        authService.AuthService
	ProvinceService    provinceService.ProvinceService
	DistrictService    districtService.DistrictService
	SubDistrictService subDistrictService.SubDistrictService
}

func initializeServices(repos *Repositories, redisClient pkg.RedisClient) *Services {
	return &Services{
		UserService: userService.NewUserService(
			userService.NewRead(repos.UserRepository, redisClient),
			userService.NewCreate(repos.UserRepository, redisClient),
			userService.NewUpdate(repos.UserRepository, repos.UserHistoryRepository, redisClient),
			userService.NewDelete(repos.UserRepository, redisClient),
			userService.NewTransaction(repos.UserRepository, redisClient),
		),
		AuthService: authService.NewAuthService(repos.UserRepository, repos.UserHistoryRepository, redisClient),
		ProvinceService: provinceService.NewProvinceService(
			provinceService.NewRead(repos.ProvinceRepository, redisClient),
		),
		DistrictService: districtService.NewDistrictService(
			districtService.NewRead(repos.DistrictRepository, redisClient),
		),
		SubDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewRead(repos.SubDistrictRepository, redisClient),
		),
	}
}
