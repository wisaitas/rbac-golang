package province

type ProvinceService interface {
	Get
	Post
}

type provinceService struct {
	Get
	Post
}

func NewProvinceService(
	get Get,
	post Post,
) ProvinceService {
	return &provinceService{
		Get:  get,
		Post: post,
	}
}
