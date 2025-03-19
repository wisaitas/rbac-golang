package district

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/wisaitas/rbac-golang/internal/auth-service/dtos/requests"
	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/internal/auth-service/repositories"
	"github.com/wisaitas/rbac-golang/pkg"
	"gorm.io/gorm"
)

func importCSV(req requests.ImportDistrict, districtRepository repositories.DistrictRepository) (statusCode int, err error) {
	return http.StatusOK, nil
}

func importJSON(req requests.ImportDistrict, districtRepository repositories.DistrictRepository) (statusCode int, err error) {
	open, err := req.File.Open()
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}
	defer open.Close()

	byteValue, err := io.ReadAll(open)
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}

	var districts []districtData
	if err := json.Unmarshal(byteValue, &districts); err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}

	tx := districtRepository.BeginTx()

	for _, district := range districts {
		province := models.Province{}
		if err := tx.Where("import_id = ?", district.ProvinceID).First(&province).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return http.StatusNotFound, pkg.Error(err)
			}

			tx.Rollback()
			return http.StatusInternalServerError, pkg.Error(err)
		}

		districtModel := district.ToModel()
		districtModel.ProvinceID = province.ID
		if err := tx.Create(&districtModel).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, pkg.Error(err)
		}
	}

	tx.Commit()

	return http.StatusOK, nil
}
