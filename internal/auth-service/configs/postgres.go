package configs

import (
	"fmt"
	"log"

	"github.com/wisaitas/rbac-golang/internal/auth-service/env"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/scripts"
	"github.com/wisaitas/rbac-golang/pkg"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		env.ENV.DB_HOST,
		env.ENV.DB_USER,
		env.ENV.DB_PASSWORD,
		env.ENV.DB_NAME,
		env.ENV.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("database connected successfully")

	if err := autoMigrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	if err := scripts.SeedData(db, scripts.Path{
		ProvincePath:       env.ENV.PROVINCE_FILE_PATH,
		DistrictPath:       env.ENV.DISTRICT_FILE_PATH,
		SubDistrictPath:    env.ENV.SUB_DISTRICT_FILE_PATH,
		PermissionPath:     env.ENV.PERMISSION_FILE_PATH,
		RolePath:           env.ENV.ROLE_FILE_PATH,
		RolePermissionPath: env.ENV.ROLE_PERMISSION_FILE_PATH,
		UserPath:           env.ENV.USER_FILE_PATH,
		UserRolePath:       env.ENV.USER_ROLE_FILE_PATH,
	}); err != nil {
		log.Fatalf("failed to seed database: %v", err)
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.Permission{},
		&models.User{},
		&models.Role{},
		&models.UserRole{},
		&models.RolePermission{},
		&models.UserHistory{},
		&models.Province{},
		&models.District{},
		&models.SubDistrict{},
		&models.Address{},
	); err != nil {
		return pkg.Error(err)
	}

	log.Println("database migrated successfully")

	return nil
}
