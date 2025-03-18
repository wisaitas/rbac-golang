package initial

import (
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
	"gorm.io/gorm"
)

type Repositories struct {
	AddressRepository        repositories.AddressRepository
	DistrictRepository       repositories.DistrictRepository
	ProvinceRepository       repositories.ProvinceRepository
	RoleRepository           repositories.RoleRepository
	RolePermissionRepository repositories.RolePermissionRepository
	SubDistrictRepository    repositories.SubDistrictRepository
	UserRepository           repositories.UserRepository
	UserHistoryRepository    repositories.UserHistoryRepository
	UserRoleRepository       repositories.UserRoleRepository
}

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		AddressRepository:        repositories.NewAddressRepository(db, pkg.NewBaseRepository[models.Address](db)),
		DistrictRepository:       repositories.NewDistrictRepository(db, pkg.NewBaseRepository[models.District](db)),
		ProvinceRepository:       repositories.NewProvinceRepository(db, pkg.NewBaseRepository[models.Province](db)),
		RoleRepository:           repositories.NewRoleRepository(db, pkg.NewBaseRepository[models.Role](db)),
		RolePermissionRepository: repositories.NewRolePermissionRepository(db, pkg.NewBaseRepository[models.RolePermission](db)),
		SubDistrictRepository:    repositories.NewSubDistrictRepository(db, pkg.NewBaseRepository[models.SubDistrict](db)),
		UserRepository:           repositories.NewUserRepository(db, pkg.NewBaseRepository[models.User](db)),
		UserHistoryRepository:    repositories.NewUserHistoryRepository(db, pkg.NewBaseRepository[models.UserHistory](db)),
		UserRoleRepository:       repositories.NewUserRoleRepository(db, pkg.NewBaseRepository[models.UserRole](db)),
	}
}
