package role

type RoleService interface {
	Create
	Read
}

type roleService struct {
	Create
	Read
}

func NewRoleService(
	create Create,
	read Read,
) RoleService {
	return &roleService{
		Create: create,
		Read:   read,
	}
}
