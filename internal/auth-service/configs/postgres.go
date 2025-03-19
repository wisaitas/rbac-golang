package configs

import (
	"fmt"
	"log"

	"github.com/wisaitas/rbac-golang/internal/auth-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		ENV.DB_HOST,
		ENV.DB_USER,
		ENV.DB_PASSWORD,
		ENV.DB_NAME,
		ENV.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := autoMigrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("database connected successfully")
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
		return fmt.Errorf("error migrating database: %w", err)
	}

	log.Println("database migrated successfully")

	return nil
}
