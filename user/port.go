package user

import (
	"gocom/domain"
	userHandler "gocom/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding rest handler user service
}

type UserRepo interface {
	Create(usr domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	Get(uId int) (*domain.User, error)
	Update(usr domain.User) (*domain.User, error)
	Delete(uId int) error
}
