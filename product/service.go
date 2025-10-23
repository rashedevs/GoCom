package product

import (
	"gocom/domain"
)

type service struct {
	prdctRepo ProductRepo
}

func NewService(prdctRepo ProductRepo) Service {
	return &service{
		prdctRepo: prdctRepo,
	}
}

func (svc *service) Create(p domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Create(p)
}

func (svc *service) Get(pId int) (*domain.Product, error) {
	return svc.prdctRepo.Get(pId)
}

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.prdctRepo.List(page, limit)
}

func (svc *service) Count() (int64, error) {
	return svc.prdctRepo.Count()
}

func (svc *service) Update(p domain.Product) (*domain.Product, error) {
	return svc.prdctRepo.Update(p)
}

func (svc *service) Delete(pId int) error {
	return svc.prdctRepo.Delete(pId)
}
