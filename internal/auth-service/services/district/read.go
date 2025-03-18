package district

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/queries"
	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/responses"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

type Read interface {
	GetDistricts(query queries.DistrictQuery) (resp []responses.DistrictResponse, statusCode int, err error)
}

type read struct {
	districtRepository repositories.DistrictRepository
	redisUtil          pkg.RedisClient
}

func NewRead(
	districtRepository repositories.DistrictRepository,
	redisUtil pkg.RedisClient,
) Read {
	return &read{
		districtRepository: districtRepository,
		redisUtil:          redisUtil,
	}
}

func (r *read) GetDistricts(query queries.DistrictQuery) (resp []responses.DistrictResponse, statusCode int, err error) {
	districts := []models.District{}

	cacheKey := fmt.Sprintf("get_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.ProvinceID)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return nil, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.districtRepository.GetAll(&districts, &query.PaginationQuery, map[string]interface{}{"province_id": query.ProvinceID}); err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, district := range districts {
		respDistrict := responses.DistrictResponse{}
		resp = append(resp, respDistrict.ModelToResponse(district))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}
