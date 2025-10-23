package product

import (
	"gocom/domain"
	productHandler "gocom/rest/handlers/product"
)

type Service interface {
	productHandler.Service // embedding rest handler product service
}

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(pId int) (*domain.Product, error)
	List(page, limit int64) ([]*domain.Product, error)
	Count() (int64, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(pId int) error
}
