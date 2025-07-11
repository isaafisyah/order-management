package repository

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/category/model"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]model.Category, error)
	FindByID(ctx context.Context, id int) (model.Category, error)
	Create(ctx context.Context, category model.Category) error
	Update(ctx context.Context, category model.Category) error
	Delete(ctx context.Context, category model.Category) error
}