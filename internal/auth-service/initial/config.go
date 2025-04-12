package initial

import (
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/rbac-golang/internal/auth-service/configs"
	"gorm.io/gorm"
)

type config struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func initializeConfig() *config {
	db := configs.ConnectDB()

	redis := configs.ConnectRedis()

	return &config{
		DB:    db,
		Redis: redis,
	}
}
