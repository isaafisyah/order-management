package usecase

import "github.com/isaafisyah/order-management/app/domain/model"

type UserUsecase interface {
	FindAll() ([]model.User, error)
	FindByID(id int) (model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user model.User) error
	
}