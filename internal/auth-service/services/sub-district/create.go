package sub_district

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/constants"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
	ImportSubDistricts(req requests.ImportSubDistrict) (statusCode int, err error)
}

type create struct {
	subDistrictRepository repositories.SubDistrictRepository
	redisClient           pkg.RedisClient
}

func NewCreate(
	subDistrictRepository repositories.SubDistrictRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		subDistrictRepository: subDistrictRepository,
		redisClient:           redisClient,
	}
}

func (r *create) ImportSubDistricts(req requests.ImportSubDistrict) (statusCode int, err error) {
	if req.File.Header.Get("Content-Type") == constants.ContentType.CSV {
		return importCSV(req, r.subDistrictRepository)
	} else if req.File.Header.Get("Content-Type") == constants.ContentType.JSON {
		return importJSON(req, r.subDistrictRepository)
	}

	return http.StatusOK, nil
}
