package tests

import (
	"testing"

	"github.com/isaafisyah/order-management/app/internal/product/request"
	"github.com/isaafisyah/order-management/app/internal/product/usecase"
	"github.com/isaafisyah/order-management/app/mocks/product"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductUseCaseSuccess(t *testing.T){
	logger.InitLogger("test")
	mockRepo := new(product.MockProductRepository)
	productUsecase := usecase.NewProductUsecase(mockRepo)

	req := request.CreateProductRequest{
		Name:     "test",
		CategoryId: 1,
		Price:    1,
		Stock: 100,
	}
	mockRepo.On("Create", mock.Anything).Return(nil).Once()
	err := productUsecase.Create(req)
	
	assert.Nil(t, err)
}

func TestCreateProductValidationFailed(t *testing.T){
	logger.InitLogger("test")
	mockRepo := new(product.MockProductRepository)
	productUsecase := usecase.NewProductUsecase(mockRepo)

	req := request.CreateProductRequest{
		Name:     "sa",
		CategoryId: 1,
		Price:    1,
		Stock: 100,
	}
	err := productUsecase.Create(req)
	
	assert.NotNil(t, err)
	assert.EqualError(t, err, "validation error")
}

