package province

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
)

func importCSV(req requests.ImportProvince, provinceRepository repositories.ProvinceRepository) (statusCode int, err error) {
	open, err := req.File.Open()
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}
	defer open.Close()

	csvReader := csv.NewReader(open)

	headers, err := csvReader.Read()
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}

	var provinces []provinceData
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return http.StatusBadRequest, pkg.Error(err)
		}

		provinceMap := make(map[string]string)
		for i, value := range record {
			provinceMap[headers[i]] = value
		}

		jsonData, err := json.Marshal(provinceMap)
		if err != nil {
			return http.StatusBadRequest, pkg.Error(err)
		}

		var province provinceData
		if err := json.Unmarshal(jsonData, &province); err != nil {
			return http.StatusBadRequest, pkg.Error(err)
		}

		provinces = append(provinces, province)
	}

	for _, province := range provinces {
		fmt.Println(province)
	}

	return http.StatusOK, nil
}

func importJSON(req requests.ImportProvince, provinceRepository repositories.ProvinceRepository) (statusCode int, err error) {
	open, err := req.File.Open()
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}
	defer open.Close()

	byteValue, err := io.ReadAll(open)
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}

	var provinces []provinceData
	if err := json.Unmarshal(byteValue, &provinces); err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}

	tx := provinceRepository.BeginTx()

	for _, province := range provinces {
		provinceModel := province.ToModel()
		if err := tx.Create(&provinceModel).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, pkg.Error(err)
		}
	}

	tx.Commit()

	return http.StatusOK, nil
}
