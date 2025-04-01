package initial

import (
	authService "github.com/wisaitas/rbac-golang/internal/auth-service/services/auth"
	districtService "github.com/wisaitas/rbac-golang/internal/auth-service/services/district"
	permissionService "github.com/wisaitas/rbac-golang/internal/auth-service/services/permission"
	provinceService "github.com/wisaitas/rbac-golang/internal/auth-service/services/province"
	roleService "github.com/wisaitas/rbac-golang/internal/auth-service/services/role"
	subDistrictService "github.com/wisaitas/rbac-golang/internal/auth-service/services/sub-district"
	userService "github.com/wisaitas/rbac-golang/internal/auth-service/services/user"
)

type Services struct {
	UserService        userService.UserService
	AuthService        authService.AuthService
	ProvinceService    provinceService.ProvinceService
	DistrictService    districtService.DistrictService
	SubDistrictService subDistrictService.SubDistrictService
	PermissionService  permissionService.PermissionService
	RoleService        roleService.RoleService
}

func initializeServices(repos *Repositories, utils *Utils) *Services {
	return &Services{
		UserService: userService.NewUserService(
			userService.NewRead(repos.UserRepository, utils.RedisUtil),
			userService.NewCreate(repos.UserRepository, utils.RedisUtil),
			userService.NewUpdate(repos.UserRepository, repos.UserHistoryRepository, utils.RedisUtil, utils.TransactionUtil),
			userService.NewDelete(repos.UserRepository, utils.RedisUtil),
			userService.NewTransaction(repos.UserRepository, utils.RedisUtil),
		),
		AuthService: authService.NewAuthService(
			repos.UserRepository,
			repos.UserHistoryRepository,
			repos.RoleRepository,
			utils.RedisUtil,
			utils.JWTUtil,
		),
		ProvinceService: provinceService.NewProvinceService(
			provinceService.NewRead(repos.ProvinceRepository, utils.RedisUtil),
			provinceService.NewCreate(repos.ProvinceRepository, utils.RedisUtil),
		),
		DistrictService: districtService.NewDistrictService(
			districtService.NewRead(repos.DistrictRepository, utils.RedisUtil),
			districtService.NewCreate(repos.DistrictRepository, utils.RedisUtil),
		),
		SubDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewRead(repos.SubDistrictRepository, utils.RedisUtil),
			subDistrictService.NewCreate(repos.SubDistrictRepository, utils.RedisUtil),
		),
		PermissionService: permissionService.NewPermissionService(
			permissionService.NewCreate(repos.PermissionRepository, utils.RedisUtil),
		),
		RoleService: roleService.NewRoleService(
			roleService.NewCreate(repos.RoleRepository, utils.RedisUtil),
			roleService.NewRead(repos.RoleRepository, utils.RedisUtil),
		),
	}
}
