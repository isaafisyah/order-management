package usecase

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/product/request"
	"github.com/isaafisyah/order-management/app/internal/product/response"
)

type ProductUsecase interface {
	FindAll(ctx context.Context) ([]response.ProductResponse, error)
	FindByID(id int) (response.ProductResponse, error)
	Create(req request.CreateProductRequest) error
}