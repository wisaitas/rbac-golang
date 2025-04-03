package sub_district

type SubDistrictService interface {
	Get
	Post
}

type subDistrictService struct {
	Get
	Post
}

func NewSubDistrictService(
	get Get,
	post Post,
) SubDistrictService {
	return &subDistrictService{
		Get:  get,
		Post: post,
	}
}
