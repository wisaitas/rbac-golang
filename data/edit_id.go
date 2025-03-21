package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

// Structs to match your JSON structure
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

func main() {
	// Read JSON files
	provinces := readProvinces("data/provinces.json")
	districts := readDistricts("data/districts.json")
	subDistricts := readSubDistricts("data/sub_districts.json")
	permissions := readPermissions("data/permissions.json")
	roles := readRoles("data/roles.json")
	users := readUsers("data/users.json")

	// Update provinces with UUIDs
	provinceIDToUUID := make(map[int]string)
	for i := range provinces {
		provinces[i].UUID = uuid.New().String()
		provinceIDToUUID[provinces[i].ID] = provinces[i].UUID
	}

	// Update districts with UUIDs and add province_uuid references
	districtIDToUUID := make(map[int]string)
	for i := range districts {
		districts[i].UUID = uuid.New().String()
		districts[i].ProvinceUUID = provinceIDToUUID[districts[i].ProvinceID]
		districtIDToUUID[districts[i].ID] = districts[i].UUID
	}

	// Update sub_districts with UUIDs and add district_uuid references
	for i := range subDistricts {
		subDistricts[i].UUID = uuid.New().String()
		subDistricts[i].DistrictUUID = districtIDToUUID[subDistricts[i].DistrictID]
	}

	// Update roles with UUIDs
	for i := range roles {
		roles[i].ID = uuid.New().String()
	}

	// Update permissions with UUIDs
	for i := range permissions {
		permissions[i].ID = uuid.New().String()
	}

	// Update users with UUIDs
	for i := range users {
		users[i].ID = uuid.New().String()
	}

	// Create roles-permissions relationship
	rolesPermissions := []RolePermission{}
	permissionNameToID := make(map[string]string)
	for _, permission := range permissions {
		permissionNameToID[permission.Name] = permission.ID
	}

	// Find admin role ID
	var adminRoleID string
	for _, role := range roles {
		if role.Name == "admin" {
			adminRoleID = role.ID
			break
		}
	}

	// Assign all permissions to admin role
	for _, permission := range permissions {
		rolesPermissions = append(rolesPermissions, RolePermission{
			RoleID:       adminRoleID,
			PermissionID: permission.ID,
		})
	}

	// Create users-roles relationship
	usersRoles := []UserRole{}
	var adminUserID string
	for _, user := range users {
		if user.Username == "admin" {
			adminUserID = user.ID
			break
		}
	}

	// Assign admin role to admin user
	usersRoles = append(usersRoles, UserRole{
		UserID: adminUserID,
		RoleID: adminRoleID,
	})

	// Write back to files
	writeToFile("data/provinces.json", provinces)
	writeToFile("data/districts.json", districts)
	writeToFile("data/sub_districts.json", subDistricts)
	writeToFile("data/permissions.json", permissions)
	writeToFile("data/roles.json", roles)
	writeToFile("data/users.json", users)
	writeToFile("data/roles_permissions.json", rolesPermissions)
	writeToFile("data/users-roles.json", usersRoles)

	fmt.Println("Edit ID Successfully")
}

// Helper functions to read JSON files
func readProvinces(filename string) []Province {
	data, err := ioutil.ReadFile(filename)
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
	data, err := ioutil.ReadFile(filename)
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
	data, err := ioutil.ReadFile(filename)
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
	data, err := ioutil.ReadFile(filename)
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
	data, err := ioutil.ReadFile(filename)
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
	data, err := ioutil.ReadFile(filename)
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

// Helper function to write JSON to file
func writeToFile(filename string, data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling data for %s: %v\n", filename, err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile(filename, jsonData, 0644); err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		os.Exit(1)
	}
}
