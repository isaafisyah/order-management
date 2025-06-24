package tests

import (
	"testing"

	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/isaafisyah/order-management/app/internal/user/request"
	"github.com/isaafisyah/order-management/app/internal/user/usecase"
	"github.com/isaafisyah/order-management/app/mocks/user"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/isaafisyah/order-management/app/utils/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserSuccess(t *testing.T) {
	mockRepo := new(user.MockUserRepository)
	mockAuthUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, mockAuthUsecase)

	req := request.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	err := validator.ValidateStruct(&req)

	assert.NoError(t, err)

	mockRepo.On("FindByEmail", mock.Anything).Return(model.User{},nil)

	mockAuthUsecase.On("HashPassword", req.Password).Return("hashedPassword", nil).Once()

	mockRepo.On("Create", mock.AnythingOfType("model.User")).Return(nil).Once()

	err = userUsecase.Create(req)
	assert.Nil(t, err)
	
	mockRepo.AssertExpectations(t)
	mockAuthUsecase.AssertExpectations(t)
}

func TestValidationFailed(t *testing.T) {
	logger.InitLogger("test")
	mockRepo := new(user.MockUserRepository)
	mockAuthUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, mockAuthUsecase)

	req := request.CreateUserRequest{
		Name:     "",
		Email:    "fsdfa",
		Password: "password123",
	}

	err := userUsecase.Create(req)
	
	assert.NotNil(t, err)
}

func TestEmailExist(t *testing.T) {
	logger.InitLogger("test")
	mockRepo := new(user.MockUserRepository)
	mockAuthUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, mockAuthUsecase)

	req := request.CreateUserRequest{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	mockRepo.On("FindByEmail", mock.Anything).Return(model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "hashedPassword",
	}, nil).Once()
	
	err := userUsecase.Create(req)
	
	assert.NotNil(t, err)
	assert.EqualError(t, err, "email already exists")
}