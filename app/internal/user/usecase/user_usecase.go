package usecase

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/user/request"
	"github.com/isaafisyah/order-management/app/internal/user/response"
)

type UserUsecase interface {
	FindAll(ctx context.Context) ([]response.UserResponse, error)
	FindByID(ctx context.Context,id int) (response.UserResponse, error)
	FindByEmail(ctx context.Context, email string) (response.UserResponse, error)
	Create(ctx context.Context ,req request.CreateUserRequest) error
	Login(ctx context.Context ,email, password string) (*string, error)
}