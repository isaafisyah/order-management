package usecase

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/category/model"
	"github.com/isaafisyah/order-management/app/internal/category/request"
)

type CategoryUsecase interface {
	FindAll(ctx context.Context) ([]model.Category, error)
	FindByID(ctx context.Context, id int) (model.Category, error)
	Create(ctx context.Context, req request.CreateCategoryRequest) error
	Update(ctx context.Context, req request.UpdateCategoryRequest) error
	Delete(ctx context.Context, id int) error
}