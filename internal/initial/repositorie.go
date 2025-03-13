package initial

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"gorm.io/gorm"
)

type Repositories struct {
	AddressRepository    repositories.AddressRepository
	CategoryRepository    repositories.CategoryRepository
	DistrictRepository    repositories.DistrictRepository
	OrderRepository    repositories.OrderRepository
	OrderItemRepository    repositories.OrderItemRepository
	PermissionRepository    repositories.PermissionRepository
	ProductRepository    repositories.ProductRepository
	ProductImageRepository    repositories.ProductImageRepository
	ProvinceRepository    repositories.ProvinceRepository
	RoleRepository    repositories.RoleRepository
	RolePermissionRepository    repositories.RolePermissionRepository
	ShippingAddressRepository    repositories.ShippingAddressRepository
	SubDistrictRepository    repositories.SubDistrictRepository
	UserRepository    repositories.UserRepository
	UserHistoryRepository    repositories.UserHistoryRepository
	UserRoleRepository    repositories.UserRoleRepository
	WishlistRepository    repositories.WishlistRepository
}

func initializeRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		AddressRepository:    repositories.NewAddressRepository(db, repositories.NewBaseRepository[models.Address](db)),
		CategoryRepository:    repositories.NewCategoryRepository(db, repositories.NewBaseRepository[models.Category](db)),
		DistrictRepository:    repositories.NewDistrictRepository(db, repositories.NewBaseRepository[models.District](db)),
		OrderRepository:    repositories.NewOrderRepository(db, repositories.NewBaseRepository[models.Order](db)),
		OrderItemRepository:    repositories.NewOrderItemRepository(db, repositories.NewBaseRepository[models.OrderItem](db)),
		PermissionRepository:    repositories.NewPermissionRepository(db, repositories.NewBaseRepository[models.Permission](db)),
		ProductRepository:    repositories.NewProductRepository(db, repositories.NewBaseRepository[models.Product](db)),
		ProductImageRepository:    repositories.NewProductImageRepository(db, repositories.NewBaseRepository[models.ProductImage](db)),
		ProvinceRepository:    repositories.NewProvinceRepository(db, repositories.NewBaseRepository[models.Province](db)),
		RoleRepository:    repositories.NewRoleRepository(db, repositories.NewBaseRepository[models.Role](db)),
		RolePermissionRepository:    repositories.NewRolePermissionRepository(db, repositories.NewBaseRepository[models.RolePermission](db)),
		ShippingAddressRepository:    repositories.NewShippingAddressRepository(db, repositories.NewBaseRepository[models.ShippingAddress](db)),
		SubDistrictRepository:    repositories.NewSubDistrictRepository(db, repositories.NewBaseRepository[models.SubDistrict](db)),
		UserRepository:    repositories.NewUserRepository(db, repositories.NewBaseRepository[models.User](db)),
		UserHistoryRepository:    repositories.NewUserHistoryRepository(db, repositories.NewBaseRepository[models.UserHistory](db)),
		UserRoleRepository:    repositories.NewUserRoleRepository(db, repositories.NewBaseRepository[models.UserRole](db)),
		WishlistRepository:    repositories.NewWishlistRepository(db, repositories.NewBaseRepository[models.Wishlist](db)),
	}
}
