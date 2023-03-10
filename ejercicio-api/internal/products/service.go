package products

import (
	"errors"
	"go-web-api/internal/domain"
)

var (
	ErrProductCodeAlreadyExist = errors.New("Code value already exists")
)

type Service interface {
	// read
	GetAllProducts() (products []domain.Product, err error)
	GetProductById(id int) (domain.Product, error)
	GetProductsMoreExpensiveThan(price float64) []domain.Product

	// write
	CreateProduct(name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error)
	Update(id int, name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	rp Repository
}

func NewService(rp Repository) Service {
	return &service{rp}
}

func (service *service) GetAllProducts() (products []domain.Product, err error) {
	return service.rp.GetAllProducts()
}

func (service *service) GetProductById(id int) (domain.Product, error) {
	return service.rp.GetProductById(id)
}

func (service *service) GetProductsMoreExpensiveThan(price float64) []domain.Product {
	return service.rp.GetProductsMoreExpensiveThan(price)
}

func (service *service) CreateProduct(name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error) {
	newProduct := domain.Product{
		Name:         name,
		Quantity:     quantity,
		Code_value:   code_value,
		Is_published: is_published,
		Expiration:   expiration,
		Price:        price,
	}
	idCreated, err := service.rp.CreateProduct(newProduct)
	if err != nil {
		return domain.Product{}, err
	}

	newProduct.ID = idCreated

	return newProduct, nil
}

func (service *service) Update(id int, name string, quantity int, code_value string, is_published bool, expiration string, price float64) (domain.Product, error) {
	newProduct := domain.Product{
		Name:         name,
		Quantity:     quantity,
		Code_value:   code_value,
		Is_published: is_published,
		Expiration:   expiration,
		Price:        price,
	}
	product, err := service.rp.Update(id, newProduct)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (service *service) Delete(id int) error {
	return service.rp.Delete(id)
}
