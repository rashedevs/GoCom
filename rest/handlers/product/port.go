package product

import "gocom/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(pId int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(pId int) error
}
