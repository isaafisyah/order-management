package tests

import (
	"testing"

	"github.com/isaafisyah/order-management/app/internal/product/model"
	"github.com/isaafisyah/order-management/app/internal/product/usecase"
	"github.com/isaafisyah/order-management/app/mocks/product"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestDeleteProductUseCaseSuccess(t *testing.T){
	productRepo := new(product.MockProductRepository)
	productUsecase := usecase.NewProductUsecase(productRepo)

	productRepo.On("Delete", mock.Anything).Return(nil).Once()
	err := productUsecase.Delete(1)
	
	assert.Nil(t, err)
}

func TestDeleteProductNotFoundFailed(t *testing.T){
	logger.InitLogger("test")
	productRepo := new(product.MockProductRepository)
	productUsecase := usecase.NewProductUsecase(productRepo)

	productRepo.On("FindByID", mock.Anything).Return(model.Product{}, nil).Once()
	err := productUsecase.Delete(1)
	
	assert.NotNil(t, err)
}