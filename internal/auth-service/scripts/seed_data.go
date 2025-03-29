package scripts

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/wisaitas/rbac-golang/internal/auth-service/models"
	"github.com/wisaitas/rbac-golang/pkg"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB, path Path) error {
	GenerateSeedID(path)

	if err := seedProvinces(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedDistricts(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedSubDistricts(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedPermissions(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedRoles(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedRolePermissions(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedUsers(db, path); err != nil {
		return pkg.Error(err)
	}

	if err := seedUserRoles(db, path); err != nil {
		return pkg.Error(err)
	}

	log.Println("database seeded successfully")
	return nil
}

type provinceData struct {
	UUID   uuid.UUID `json:"uuid"`
	NameTH string    `json:"name_th"`
	NameEN string    `json:"name_en"`
}

func (r *provinceData) ToModel() models.Province {
	return models.Province{
		BaseModel: pkg.BaseModel{
			ID: r.UUID,
		},
		NameTH: r.NameTH,
		NameEN: r.NameEN,
	}
}

func seedProvinces(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.Province{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.ProvincePath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []provinceData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var provinces []models.Province
		for _, data := range datas {
			provinces = append(provinces, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&provinces, 100).Error)
	}

	return nil
}

type districtData struct {
	UUID   uuid.UUID `json:"uuid"`
	NameTH string    `json:"name_th"`
	NameEN string    `json:"name_en"`

	ProvinceUUID uuid.UUID `json:"province_uuid"`
}

func (r *districtData) ToModel() models.District {
	return models.District{
		BaseModel: pkg.BaseModel{
			ID: r.UUID,
		},
		NameTH:     r.NameTH,
		NameEN:     r.NameEN,
		ProvinceID: r.ProvinceUUID,
	}
}

func seedDistricts(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.District{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.DistrictPath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []districtData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var districts []models.District
		for _, data := range datas {
			districts = append(districts, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&districts, 100).Error)
	}

	return nil
}

type subDistrictData struct {
	UUID    uuid.UUID `json:"uuid"`
	NameTH  string    `json:"name_th"`
	NameEN  string    `json:"name_en"`
	ZipCode int       `json:"zip_code"`

	DistrictUUID uuid.UUID `json:"district_uuid"`
}

func (r *subDistrictData) ToModel() models.SubDistrict {
	return models.SubDistrict{
		BaseModel: pkg.BaseModel{
			ID: r.UUID,
		},
		NameTH:     r.NameTH,
		NameEN:     r.NameEN,
		PostalCode: strconv.Itoa(r.ZipCode),
		DistrictID: r.DistrictUUID,
	}
}

func seedSubDistricts(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.SubDistrict{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.SubDistrictPath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []subDistrictData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var subDistricts []models.SubDistrict
		for _, data := range datas {
			subDistricts = append(subDistricts, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&subDistricts, 100).Error)
	}

	return nil
}

type permissionData struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func (r *permissionData) ToModel() models.Permission {
	return models.Permission{
		BaseModel: pkg.BaseModel{
			ID: r.ID,
		},
		Name:        r.Name,
		Description: r.Description,
	}
}

func seedPermissions(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.Permission{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.PermissionPath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []permissionData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var permissions []models.Permission
		for _, data := range datas {
			permissions = append(permissions, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&permissions, 100).Error)
	}
	return nil
}

type roleData struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description *string     `json:"description"`
	Permissions []uuid.UUID `json:"permissions"`
}

func (r *roleData) ToModel() models.Role {
	return models.Role{
		BaseModel: pkg.BaseModel{
			ID: r.ID,
		},
		Name:        r.Name,
		Description: r.Description,
	}
}

func seedRoles(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.Role{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.RolePath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []roleData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var roles []models.Role
		for _, data := range datas {
			roles = append(roles, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&roles, 100).Error)
	}
	return nil
}

type rolePermissionData struct {
	RoleID       uuid.UUID `json:"role_id"`
	PermissionID uuid.UUID `json:"permission_id"`
}

func (r *rolePermissionData) ToModel() models.RolePermission {
	return models.RolePermission{
		RoleID:       r.RoleID,
		PermissionID: r.PermissionID,
	}
}

func seedRolePermissions(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.RolePermission{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.RolePermissionPath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []rolePermissionData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var rolePermissions []models.RolePermission
		for _, data := range datas {
			rolePermissions = append(rolePermissions, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&rolePermissions, 100).Error)
	}
	return nil
}

type userData struct {
	ID        uuid.UUID   `json:"id"`
	Username  string      `json:"username"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	BirthDate time.Time   `json:"birth_date"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Roles     []uuid.UUID `json:"roles"`
}

func (r *userData) ToModel() models.User {
	return models.User{
		BaseModel: pkg.BaseModel{ID: r.ID},
		Username:  r.Username,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		BirthDate: r.BirthDate,
		Email:     r.Email,
		Password:  r.Password,
	}
}

func seedUsers(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.UserPath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []userData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var users []models.User
		for _, data := range datas {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
			if err != nil {
				return pkg.Error(err)
			}

			data.Password = string(hashedPassword)
			users = append(users, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&users, 100).Error)
	}
	return nil
}

type userRoleData struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}

func (r *userRoleData) ToModel() models.UserRole {
	return models.UserRole{
		UserID: r.UserID,
		RoleID: r.RoleID,
	}
}

func seedUserRoles(db *gorm.DB, path Path) error {
	var count int64
	if err := db.Model(&models.UserRole{}).Count(&count).Error; err != nil {
		return pkg.Error(err)
	}

	if count == 0 {
		file, err := os.Open(path.UserRolePath)
		if err != nil {
			return pkg.Error(err)
		}
		defer file.Close()

		byteData, err := io.ReadAll(file)
		if err != nil {
			return pkg.Error(err)
		}

		var datas []userRoleData
		if err := json.Unmarshal(byteData, &datas); err != nil {
			return pkg.Error(err)
		}

		var userRoles []models.UserRole
		for _, data := range datas {
			userRoles = append(userRoles, data.ToModel())
		}

		return pkg.Error(db.CreateInBatches(&userRoles, 100).Error)
	}
	return nil
}
