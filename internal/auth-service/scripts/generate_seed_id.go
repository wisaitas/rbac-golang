package scripts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type Province struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid"`
	NameTH      string `json:"name_th"`
	NameEN      string `json:"name_en"`
	GeographyID int    `json:"geography_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   any    `json:"deleted_at"`
}

type District struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	ProvinceID   int    `json:"province_id"`
	ProvinceUUID string `json:"province_uuid,omitempty"`
	NameTH       string `json:"name_th"`
	NameEN       string `json:"name_en"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    any    `json:"deleted_at"`
}

type SubDistrict struct {
	ID           int    `json:"id"`
	UUID         string `json:"uuid"`
	DistrictID   int    `json:"district_id"`
	DistrictUUID string `json:"district_uuid,omitempty"`
	ZipCode      int    `json:"zip_code"`
	NameTH       string `json:"name_th"`
	NameEN       string `json:"name_en"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    any    `json:"deleted_at"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Permission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RolePermission struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

type UserRole struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

func generateSeedID(path Path) {
	provinces := readProvinces(path.ProvincePath)
	districts := readDistricts(path.DistrictPath)
	subDistricts := readSubDistricts(path.SubDistrictPath)
	permissions := readPermissions(path.PermissionPath)
	roles := readRoles(path.RolePath)
	rolesPermissions := readRolesPermissions(path.RolePermissionPath)
	users := readUsers(path.UserPath)
	usersRoles := readUsersRoles(path.UserRolePath)

	provinceIDToUUID := make(map[int]string)
	for i := range provinces {
		provinces[i].UUID = uuid.New().String()
		provinceIDToUUID[provinces[i].ID] = provinces[i].UUID
	}

	districtIDToUUID := make(map[int]string)
	for i := range districts {
		districts[i].UUID = uuid.New().String()
		districts[i].ProvinceUUID = provinceIDToUUID[districts[i].ProvinceID]
		districtIDToUUID[districts[i].ID] = districts[i].UUID
	}

	subDistrictIDToUUID := make(map[int]string)
	for i := range subDistricts {
		subDistricts[i].UUID = uuid.New().String()
		subDistricts[i].DistrictUUID = districtIDToUUID[subDistricts[i].DistrictID]
		subDistrictIDToUUID[subDistricts[i].ID] = subDistricts[i].UUID
	}

	roleIDMap := make(map[string]string)
	for i := range roles {
		oldID := roles[i].ID
		roles[i].ID = uuid.New().String()
		roleIDMap[oldID] = roles[i].ID
	}

	permissionIDMap := make(map[string]string)
	for i := range permissions {
		oldID := permissions[i].ID
		permissions[i].ID = uuid.New().String()
		permissionIDMap[oldID] = permissions[i].ID
	}

	for i := range rolesPermissions {
		rolesPermissions[i].RoleID = roleIDMap[rolesPermissions[i].RoleID]
		rolesPermissions[i].PermissionID = permissionIDMap[rolesPermissions[i].PermissionID]
	}

	userIDMap := make(map[string]string)
	for i := range users {
		oldID := users[i].ID
		users[i].ID = uuid.New().String()
		userIDMap[oldID] = users[i].ID
	}

	for i := range usersRoles {
		usersRoles[i].UserID = userIDMap[usersRoles[i].UserID]
		usersRoles[i].RoleID = roleIDMap[usersRoles[i].RoleID]
	}

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

func readRolesPermissions(filename string) []RolePermission {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}
	var rolesPermissions []RolePermission
	if err := json.Unmarshal(data, &rolesPermissions); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return rolesPermissions
}

func readUsersRoles(filename string) []UserRole {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		os.Exit(1)
	}

	var usersRoles []UserRole
	if err := json.Unmarshal(data, &usersRoles); err != nil {
		fmt.Printf("Error unmarshaling %s: %v\n", filename, err)
		os.Exit(1)
	}
	return usersRoles
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
