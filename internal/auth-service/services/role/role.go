package role

type RoleService interface {
	Create
}

type roleService struct {
	Create
}

func NewRoleService(
	create Create,
) RoleService {
	return &roleService{
		Create: create,
	}
}
