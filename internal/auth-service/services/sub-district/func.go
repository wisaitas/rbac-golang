package sub_district

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

func importCSV(req requests.ImportSubDistrict, subDistrictRepository repositories.SubDistrictRepository) (statusCode int, err error) {
	return http.StatusOK, nil
}

func importJSON(req requests.ImportSubDistrict, subDistrictRepository repositories.SubDistrictRepository) (statusCode int, err error) {
	open, err := req.File.Open()
	if err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}
	defer open.Close()

	byteValue, err := io.ReadAll(open)
	if err != nil {

		return http.StatusBadRequest, pkg.Error(err)
	}

	var subDistricts []subDistrictData
	if err := json.Unmarshal(byteValue, &subDistricts); err != nil {
		return http.StatusBadRequest, pkg.Error(err)
	}

	tx := subDistrictRepository.BeginTx()

	for _, subDistrict := range subDistricts {
		district := models.District{}
		if err := tx.Where("import_id = ?", subDistrict.DistrictID).First(&district).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return http.StatusNotFound, pkg.Error(err)
			}

			tx.Rollback()
			return http.StatusInternalServerError, pkg.Error(err)
		}

		subDistrictModel := subDistrict.ToModel()
		subDistrictModel.DistrictID = district.ID
		if err := tx.Create(&subDistrictModel).Error; err != nil {
			tx.Rollback()
			return http.StatusInternalServerError, pkg.Error(err)
		}
	}

	tx.Commit()

	return http.StatusOK, nil
}
