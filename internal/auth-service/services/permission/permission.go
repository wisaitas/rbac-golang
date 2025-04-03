package permission

type PermissionService interface {
	Post
	Get
}

type permissionService struct {
	Post
	Get
}

func NewPermissionService(
	post Post,
	get Get,
) PermissionService {
	return &permissionService{
		Post: post,
		Get:  get,
	}
}
