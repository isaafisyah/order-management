package repository

import "github.com/isaafisyah/order-management/app/internal/user/model"

type UserRepository interface{
	FindAll() ([]model.User, error)
	FindByID(id int) (model.User, error)
	FindByEmail(email string) (model.User, error)
	Create(user model.User) error
}