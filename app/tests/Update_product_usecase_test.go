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

func TestUpdateProductUseCaseSuccess(t *testing.T){
	productRepo := new(product.MockProductRepository)
	productUsecase := usecase.NewProductUsecase(productRepo)

	req := request.UpdateProductRequest{
		Name:     "test",
		CategoryId: 1,
		Price:    1,
		Stock: 100,
	}
	productRepo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()
	err := productUsecase.Update(1, req)
	
	assert.Nil(t, err)
}

func TestUpdateProductValidationFailed(t *testing.T){
	logger.InitLogger("test")
	mockRepo := new(product.MockProductRepository)
	productUsecase := usecase.NewProductUsecase(mockRepo)

	req := request.UpdateProductRequest{
		Name:     "sa",
		CategoryId: 1,
		Price:    1,
		Stock: 100,
	}
	err := productUsecase.Update(1, req)
	
	assert.NotNil(t, err)
	assert.EqualError(t, err, "validation error")
}
