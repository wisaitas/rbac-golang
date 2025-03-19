package province

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/constants"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
	ImportProvinces(req requests.ImportProvince) (statusCode int, err error)
}

type create struct {
	provinceRepository repositories.ProvinceRepository
	redisClient        pkg.RedisClient
}

func NewCreate(
	provinceRepository repositories.ProvinceRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		provinceRepository: provinceRepository,
		redisClient:        redisClient,
	}
}

func (r *create) ImportProvinces(req requests.ImportProvince) (statusCode int, err error) {
	if req.File.Header.Get("Content-Type") == constants.ContentType.CSV {
		return importCSV(req, r.provinceRepository)
	} else if req.File.Header.Get("Content-Type") == constants.ContentType.JSON {
		return importJSON(req, r.provinceRepository)
	}

	return http.StatusOK, nil
}
