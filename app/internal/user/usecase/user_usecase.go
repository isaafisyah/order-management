package usecase

import (
	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/isaafisyah/order-management/app/internal/user/response"
)



type UserUsecase interface {
	FindAll() ([]response.UserResponse, error)
	FindByID(id int) (response.UserResponse, error)
	FindByEmail(email string) (response.UserResponse, error)
	Create(user model.User) error
	Login(email, password string) (*string, error)
	
}