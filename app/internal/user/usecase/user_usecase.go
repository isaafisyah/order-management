package usecase

import (
	"github.com/isaafisyah/order-management/app/internal/user/request"
	"github.com/isaafisyah/order-management/app/internal/user/response"
)



type UserUsecase interface {
	FindAll() ([]response.UserResponse, error)
	FindByID(id int) (response.UserResponse, error)
	FindByEmail(email string) (response.UserResponse, error)
	Create(req request.CreateUserRequest) error
	Login(email, password string) (*string, error)
}