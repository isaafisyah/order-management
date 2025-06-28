package usecase

import (
	"context"
	"errors"

	"github.com/isaafisyah/order-management/app/internal/product/repository"
	"github.com/isaafisyah/order-management/app/internal/product/request"
	"github.com/isaafisyah/order-management/app/internal/product/response"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/isaafisyah/order-management/app/utils/validator"
)

type ProductUsecaseImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{
		ProductRepository: productRepository,
	}
}

func (u *ProductUsecaseImpl) FindAll(ctx context.Context) ([]response.ProductResponse, error) {
	products, _ := u.ProductRepository.FindAll(ctx)
	return response.ToProductResponses(products), nil
}

func (u *ProductUsecaseImpl) FindByID(id int) (response.ProductResponse, error)  {
	product, err := u.ProductRepository.FindByID(id)
	return response.ToProductResponse(product), err
}

func (u *ProductUsecaseImpl) Create(req request.CreateProductRequest) error {
	if err := validator.ValidateStruct(&req); err != nil {
		err = errors.New("validation error")
		logger.Log.WithField("Module", "ProductService").WithError(err).Error("Failed to create product")
		return err
	}
	product := request.ToProduct(req)
	
	return u.ProductRepository.Create(product)
}