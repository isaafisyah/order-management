package tests

import (
	"errors"
	"testing"

	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/isaafisyah/order-management/app/internal/user/usecase"
	"github.com/isaafisyah/order-management/app/mocks/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginSuccess(t *testing.T)  {
	mockRepo := new(user.MockUserRepository)
	authUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, authUsecase)

	mockRepo.On("FindByEmail", mock.Anything).Return(model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "hashedPassword",
	}, nil).Once()
	
	authUsecase.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(true).Once()

	authUsecase.On("GenerateToken", mock.Anything).Return("token", nil).Once()
	
	token, err := userUsecase.Login("john.doe@example.com", "password123")
	
	assert.Nil(t, err)
	assert.NotNil(t, token)
}

func TestUserNotFoundFailed(t *testing.T)  {
	mockRepo := new(user.MockUserRepository)
	authUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, authUsecase)

	mockRepo.On("FindByEmail", mock.Anything).Return(model.User{}, nil).Once()
	
	token, err := userUsecase.Login("john.doe@example.com", "password123")
	
	assert.NotNil(t, err)
	assert.Nil(t, token)
	assert.EqualError(t, err, "user not found")
}

func TestInvalidPasswordFailed(t *testing.T)  {
	mockRepo := new(user.MockUserRepository)
	authUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, authUsecase)

	mockRepo.On("FindByEmail", mock.Anything).Return(model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "hashedPassword",
	}, nil).Once()
	
	authUsecase.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(false).Once()
	
	token, err := userUsecase.Login("john.doe@example.com", "password123")
	
	assert.NotNil(t, err)
	assert.Nil(t, token)
	assert.EqualError(t, err, "invalid password")
}

func TestGenerateTokenFailed(t *testing.T)  {
	mockRepo := new(user.MockUserRepository)
	authUsecase := new(user.MockAuthUsecase)
	userUsecase := usecase.NewUserUsecase(mockRepo, authUsecase)

	mockRepo.On("FindByEmail", mock.Anything).Return(model.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "hashedPassword",
	}, nil).Once()
	
	authUsecase.On("CheckPasswordHash", mock.Anything, mock.Anything).Return(true).Once()

	error := errors.New("error generate token")
	authUsecase.On("GenerateToken", mock.Anything).Return("", error).Once()
	
	token, err := userUsecase.Login("john.doe@example.com", "password123")
	
	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.EqualError(t, err, "failed to generate token")
}