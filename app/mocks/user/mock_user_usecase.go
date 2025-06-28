package user

import (
	"context"
	"time"

	"github.com/isaafisyah/order-management/app/internal/user/request"
	"github.com/isaafisyah/order-management/app/internal/user/response"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Login(ctx context.Context, email string, password string) (*string, error) {
	args := m.Called(ctx, email, password)
	str := args.Get(0).(*string)
	return str, args.Error(1)
}

func (m *MockUserUsecase) FindAll(ctx context.Context) ([]response.UserResponse, error) {
	args := m.Called(ctx)
	select {
    case <-time.After(7 * time.Second):
        return args.Get(0).([]response.UserResponse), args.Error(1)
    case <-ctx.Done():
        return []response.UserResponse{}, ctx.Err()
    }
}

func (m *MockUserUsecase) FindByID(ctx context.Context, id int) (response.UserResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(response.UserResponse), args.Error(1)
}

func (m *MockUserUsecase) FindByEmail(ctx context.Context, email string) (response.UserResponse, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(response.UserResponse), args.Error(1)
}

func (m *MockUserUsecase) Create(ctx context.Context, req request.CreateUserRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}