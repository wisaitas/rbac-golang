package scripts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type Province struct {
	ID          int       `json:"id"`
	UUID        uuid.UUID `json:"uuid"`
	NameTH      string    `json:"name_th"`
	NameEN      string    `json:"name_en"`
	GeographyID int       `json:"geography_id"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	DeletedAt   any       `json:"deleted_at"`
}

type District struct {
	ID           int       `json:"id"`
	UUID         uuid.UUID `json:"uuid"`
	ProvinceID   int       `json:"province_id"`
	ProvinceUUID uuid.UUID `json:"province_uuid,omitempty"`
	NameTH       string    `json:"name_th"`
	NameEN       string    `json:"name_en"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	DeletedAt    any       `json:"deleted_at"`
}

type SubDistrict struct {
	ID           int       `json:"id"`
	UUID         uuid.UUID `json:"uuid"`
	DistrictID   int       `json:"district_id"`
	DistrictUUID uuid.UUID `json:"district_uuid,omitempty"`
	ZipCode      int       `json:"zip_code"`
	NameTH       string    `json:"name_th"`
	NameEN       string    `json:"name_en"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	DeletedAt    any       `json:"deleted_at"`
}

type Role struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Permission struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate string    `json:"birth_date"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type RolePermission struct {
	RoleID       uuid.UUID `json:"role_id"`
	PermissionID uuid.UUID `json:"permission_id"`
}

type UserRole struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
}

func GenerateSeedID(path Path) {
	provinces := readProvinces(path.ProvincePath)
	districts := readDistricts(path.DistrictPath)
	subDistricts := readSubDistricts(path.SubDistrictPath)
	permissions := readPermissions(path.PermissionPath)
	roles := readRoles(path.RolePath)
	users := readUsers(path.UserPath)

	for _, province := range provinces {
		if province.UUID != uuid.Nil {
			return
		}
	}

	provinceIDToUUID := make(map[int]uuid.UUID)
	for i := range provinces {
		provinces[i].UUID = uuid.New()
		provinceIDToUUID[provinces[i].ID] = provinces[i].UUID
	}

	districtIDToUUID := make(map[int]uuid.UUID)
	for i := range districts {
		districts[i].UUID = uuid.New()
		districts[i].ProvinceUUID = provinceIDToUUID[districts[i].ProvinceID]
		districtIDToUUID[districts[i].ID] = districts[i].UUID
	}

	for i := range subDistricts {
		subDistricts[i].UUID = uuid.New()
		subDistricts[i].DistrictUUID = districtIDToUUID[subDistricts[i].DistrictID]
	}

	for i := range roles {
		roles[i].ID = uuid.New()
	}

	for i := range permissions {
		permissions[i].ID = uuid.New()
	}

	for i := range users {
		users[i].ID = uuid.New()
	}

	rolesPermissions := []RolePermission{}
	permissionNameToID := make(map[string]uuid.UUID)
	for _, permission := range permissions {
		permissionNameToID[permission.Name] = permission.ID
	}

	var adminRoleID uuid.UUID
	for _, role := range roles {
		if role.Name == "admin" {
			adminRoleID = role.ID
			break
		}
	}

	for _, permission := range permissions {
		rolesPermissions = append(rolesPermissions, RolePermission{
			RoleID:       adminRoleID,
			PermissionID: permission.ID,
		})
	}

	usersRoles := []UserRole{}
	var adminUserID uuid.UUID
	for _, user := range users {
		if user.Username == "admin" {
			adminUserID = user.ID
			break
		}
	}

	usersRoles = append(usersRoles, UserRole{
		UserID: adminUserID,
		RoleID: adminRoleID,
	})

	writeToFile(path.ProvincePath, provinces)
	writeToFile(path.DistrictPath, districts)
	writeToFile(path.SubDistrictPath, subDistricts)
	writeToFile(path.PermissionPath, permissions)
	writeToFile(path.RolePath, roles)
	writeToFile(path.UserPath, users)
	writeToFile(path.RolePermissionPath, rolesPermissions)
	writeToFile(path.UserRolePath, usersRoles)

	log.Println("generate seed ID successfully")
}

func readProvinces(filename string) []Province {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var provinces []Province
	if err := json.Unmarshal(data, &provinces); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return provinces
}

func readDistricts(filename string) []District {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var districts []District
	if err := json.Unmarshal(data, &districts); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return districts
}

func readSubDistricts(filename string) []SubDistrict {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var subDistricts []SubDistrict
	if err := json.Unmarshal(data, &subDistricts); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return subDistricts
}

func readPermissions(filename string) []Permission {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var permissions []Permission
	if err := json.Unmarshal(data, &permissions); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return permissions
}

func readRoles(filename string) []Role {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var roles []Role
	if err := json.Unmarshal(data, &roles); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return roles
}

func readUsers(filename string) []User {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return users
}

func writeToFile(filename string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data for %s: %v\n", filename, err)
		os.Exit(1)
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		os.Exit(1)
	}
}
