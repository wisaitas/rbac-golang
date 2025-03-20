package permission

type PermissionService interface {
	Create
}

type permissionService struct {
	Create
}

func NewPermissionService(
	create Create,
) PermissionService {
	return &permissionService{
		Create: create,
	}
}
