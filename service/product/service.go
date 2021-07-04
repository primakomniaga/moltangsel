package product

//interface to Resource
type Resource interface {
	CreateProduct(*Product) error
	Products() ([]Product, error)
	Product(int) (*Product, error)
	EditProduct(*Product) error
	DeleteProduct(int) error
	UpdateImage(int, []string) error
	ProductLimit(int, int) ([]Product, error)
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

func (s *Service) Products() ([]Product, error) {
	products, err := s.resource.Products()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *Service) Product(idProdut int) (*Product, error) {
	product, err := s.resource.Product(idProdut)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) EditProduct(model *Product) error {
	if err := s.resource.EditProduct(model); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteProduct(idproduct int) error {
	if err := s.resource.DeleteProduct(idproduct); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateImage(idProduct int, image []string) error {
	if err := s.resource.UpdateImage(idProduct, image); err != nil {
		return err
	}
	return nil
}

func (s *Service) ProductLimit(page, limit int) ([]Product, error) {
	products, err := s.resource.ProductLimit(page, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}
