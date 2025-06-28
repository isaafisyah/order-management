package repository

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/user/model"
)

type UserRepository interface{
	FindAll(ctx context.Context) ([]model.User, error)
	FindByID(ctx context.Context, id int) (model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	Create(ctx context.Context, user model.User) error
}