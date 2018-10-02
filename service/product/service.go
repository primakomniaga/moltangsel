package product

//interface to Resource
type Resource interface {
	CreateProduct(*Product) error
}

//service of product
type Service struct {
	resource Resource
}

func New(resProduct Resource) *Service {
	return &Service{
		resource: resProduct,
	}
}

func (s *Service) CreateProduct(model *Product) error {
	if err := s.resource.CreateProduct(model); err != nil {
		return err
	}
	return nil
}
