package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type repositorie struct {
	AddressRepository        repositories.AddressRepository
	DistrictRepository       repositories.DistrictRepository
	ProvinceRepository       repositories.ProvinceRepository
	RoleRepository           repositories.RoleRepository
	RolePermissionRepository repositories.RolePermissionRepository
	SubDistrictRepository    repositories.SubDistrictRepository
	UserRepository           repositories.UserRepository
	UserHistoryRepository    repositories.UserHistoryRepository
	UserRoleRepository       repositories.UserRoleRepository
	PermissionRepository     repositories.PermissionRepository
}

func initializeRepositorie(config *config) *repositorie {
	return &repositorie{
		AddressRepository:        repositories.NewAddressRepository(config.DB, pkg.NewBaseRepository[models.Address](config.DB)),
		DistrictRepository:       repositories.NewDistrictRepository(config.DB, pkg.NewBaseRepository[models.District](config.DB)),
		ProvinceRepository:       repositories.NewProvinceRepository(config.DB, pkg.NewBaseRepository[models.Province](config.DB)),
		RoleRepository:           repositories.NewRoleRepository(config.DB, pkg.NewBaseRepository[models.Role](config.DB)),
		RolePermissionRepository: repositories.NewRolePermissionRepository(config.DB, pkg.NewBaseRepository[models.RolePermission](config.DB)),
		SubDistrictRepository:    repositories.NewSubDistrictRepository(config.DB, pkg.NewBaseRepository[models.SubDistrict](config.DB)),
		UserRepository:           repositories.NewUserRepository(config.DB, pkg.NewBaseRepository[models.User](config.DB)),
		UserHistoryRepository:    repositories.NewUserHistoryRepository(config.DB, pkg.NewBaseRepository[models.UserHistory](config.DB)),
		UserRoleRepository:       repositories.NewUserRoleRepository(config.DB, pkg.NewBaseRepository[models.UserRole](config.DB)),
		PermissionRepository:     repositories.NewPermissionRepository(config.DB, pkg.NewBaseRepository[models.Permission](config.DB)),
	}
}
