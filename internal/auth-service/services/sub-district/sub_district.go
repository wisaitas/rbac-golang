package sub_district

type SubDistrictService interface {
	Read
	Create
}

type subDistrictService struct {
	Read
	Create
}

func NewSubDistrictService(
	read Read,
	create Create,
) SubDistrictService {
	return &subDistrictService{
		Read:   read,
		Create: create,
	}
}
