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

type service struct {
	UserService        userService.UserService
	AuthService        authService.AuthService
	ProvinceService    provinceService.ProvinceService
	DistrictService    districtService.DistrictService
	SubDistrictService subDistrictService.SubDistrictService
	PermissionService  permissionService.PermissionService
	RoleService        roleService.RoleService
}

func initializeService(repo *repositorie, util *util) *service {
	return &service{
		UserService: userService.NewUserService(
			userService.NewGet(repo.UserRepository, util.RedisUtil),
			userService.NewPost(repo.UserRepository, util.RedisUtil),
			userService.NewUpdate(repo.UserRepository, repo.UserHistoryRepository, util.RedisUtil, util.TransactionUtil),
			userService.NewDelete(repo.UserRepository, util.RedisUtil),
			userService.NewTransaction(repo.UserRepository, util.RedisUtil),
		),
		AuthService: authService.NewAuthService(
			repo.UserRepository,
			repo.UserHistoryRepository,
			repo.RoleRepository,
			util.RedisUtil,
			util.JWTUtil,
		),
		ProvinceService: provinceService.NewProvinceService(
			provinceService.NewGet(repo.ProvinceRepository, util.RedisUtil),
			provinceService.NewPost(repo.ProvinceRepository, util.RedisUtil),
		),
		DistrictService: districtService.NewDistrictService(
			districtService.NewGet(repo.DistrictRepository, util.RedisUtil),
			districtService.NewPost(repo.DistrictRepository, util.RedisUtil),
		),
		SubDistrictService: subDistrictService.NewSubDistrictService(
			subDistrictService.NewGet(repo.SubDistrictRepository, util.RedisUtil),
			subDistrictService.NewPost(repo.SubDistrictRepository, util.RedisUtil),
		),
		PermissionService: permissionService.NewPermissionService(
			permissionService.NewPost(repo.PermissionRepository, util.RedisUtil),
			permissionService.NewGet(repo.PermissionRepository, util.RedisUtil),
		),
		RoleService: roleService.NewRoleService(
			roleService.NewPost(repo.RoleRepository, util.RedisUtil),
			roleService.NewGet(repo.RoleRepository, util.RedisUtil),
		),
	}
}
