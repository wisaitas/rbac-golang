package role

type RoleService interface {
	Post
	Get
}

type roleService struct {
	Post
	Get
}

func NewRoleService(
	post Post,
	get Get,
) RoleService {
	return &roleService{
		Post: post,
		Get:  get,
	}
}
