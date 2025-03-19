package district

import (
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/constants"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Create interface {
	ImportDistricts(req requests.ImportDistrict) (statusCode int, err error)
}

type create struct {
	districtRepository repositories.DistrictRepository
	redisClient        pkg.RedisClient
}

func NewCreate(
	districtRepository repositories.DistrictRepository,
	redisClient pkg.RedisClient,
) Create {
	return &create{
		districtRepository: districtRepository,
		redisClient:        redisClient,
	}
}

func (r *create) ImportDistricts(req requests.ImportDistrict) (statusCode int, err error) {
	if req.File.Header.Get("Content-Type") == constants.ContentType.CSV {
		return importCSV(req, r.districtRepository)
	} else if req.File.Header.Get("Content-Type") == constants.ContentType.JSON {
		return importJSON(req, r.districtRepository)
	}

	return http.StatusOK, nil
}
