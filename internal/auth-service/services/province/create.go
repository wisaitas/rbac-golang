package province

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
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
		return importCSV(req)
	} else if req.File.Header.Get("Content-Type") == constants.ContentType.JSON {
		return importJSON(req)
	}

	return http.StatusOK, nil
}

func importCSV(req requests.ImportProvince) (statusCode int, err error) {
	open, err := req.File.Open()
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer open.Close()

	csvReader := csv.NewReader(open)

	// Read header row
	headers, err := csvReader.Read()
	if err != nil {
		return http.StatusBadRequest, err
	}

	var provinces []provinceData
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return http.StatusBadRequest, err
		}

		provinceMap := make(map[string]string)
		for i, value := range record {
			provinceMap[headers[i]] = value
		}

		jsonData, err := json.Marshal(provinceMap)
		if err != nil {
			return http.StatusBadRequest, err
		}

		var province provinceData
		if err := json.Unmarshal(jsonData, &province); err != nil {
			return http.StatusBadRequest, err
		}

		provinces = append(provinces, province)
	}

	for _, province := range provinces {
		fmt.Println(province)
	}

	return http.StatusOK, nil
}

func importJSON(req requests.ImportProvince) (statusCode int, err error) {

	open, err := req.File.Open()
	if err != nil {
		return http.StatusBadRequest, err
	}
	defer open.Close()

	byteValue, err := io.ReadAll(open)
	if err != nil {
		return http.StatusBadRequest, err
	}

	var provinces []provinceData
	if err := json.Unmarshal(byteValue, &provinces); err != nil {
		return http.StatusBadRequest, err
	}

	for _, province := range provinces {
		fmt.Println(province)
	}

	return http.StatusOK, nil
}
