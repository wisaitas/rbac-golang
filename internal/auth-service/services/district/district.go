package district

type DistrictService interface {
	Read
	Create
}

type districtService struct {
	Read
	Create
}

func NewDistrictService(
	read Read,
	create Create,
) DistrictService {
	return &districtService{
		Read:   read,
		Create: create,
	}
}
