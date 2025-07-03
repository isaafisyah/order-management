package product

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/product/model"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) FindByID(id int) (model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(product model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(product model.Product) error {
	args := m.Called(product)
	return args.Error(0)
}