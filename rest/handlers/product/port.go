package product

import "gocom/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Get(pId int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Update(domain.Product) (*domain.Product, error)
	Delete(pId int) error
}
