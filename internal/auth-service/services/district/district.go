package district

type DistrictService interface {
	Get
	Post
}

type districtService struct {
	Get
	Post
}

func NewDistrictService(
	get Get,
	post Post,
) DistrictService {
	return &districtService{
		Get:  get,
		Post: post,
	}
}
