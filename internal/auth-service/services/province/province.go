package province

type ProvinceService interface {
	Read
	Create
}

type provinceService struct {
	Read
	Create
}

func NewProvinceService(
	read Read,
	create Create,
) ProvinceService {
	return &provinceService{
		Read:   read,
		Create: create,
	}
}
