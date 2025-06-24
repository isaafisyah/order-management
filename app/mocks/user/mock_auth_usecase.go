package user

import (
	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/stretchr/testify/mock"
)

type MockAuthUsecase struct {
	mock.Mock
}

func (m *MockAuthUsecase) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockAuthUsecase) CheckPasswordHash(password string, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

func (m *MockAuthUsecase) GenerateToken(user model.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}