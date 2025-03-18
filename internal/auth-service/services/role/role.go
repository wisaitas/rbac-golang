package role

type RoleService interface {
}

type roleService struct {
}

func NewRoleService() RoleService {
	return &roleService{}
}
