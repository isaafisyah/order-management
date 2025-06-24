package usecase

import (
	"errors"
	"fmt"

	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/isaafisyah/order-management/app/internal/user/repository"
	"github.com/isaafisyah/order-management/app/internal/user/request"
	"github.com/isaafisyah/order-management/app/internal/user/response"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/isaafisyah/order-management/app/utils/validator"
)


type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
	authUseCase AuthUsecase
}

func NewUserUsecase(userRepository repository.UserRepository, authUseCase AuthUsecase) UserUsecase {
	return &UserUsecaseImpl{userRepository, authUseCase}
}

func (u *UserUsecaseImpl) FindAll() ([]response.UserResponse, error) {
	users, err := u.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return response.ToUserResponses(users), nil
}

func (u *UserUsecaseImpl) FindByID(id int) (response.UserResponse, error) {
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(user), nil
}

func (u *UserUsecaseImpl) FindByEmail(email string) (response.UserResponse, error) {
	data, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return response.UserResponse{}, err
	}
	return response.ToUserResponse(data), nil
}

func (u *UserUsecaseImpl) Create(req request.CreateUserRequest) error {
	if err := validator.ValidateStruct(&req); err != nil {
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to register user")
		return fmt.Errorf("validation error: %w", err)
	}

	existEmail, _ := u.UserRepository.FindByEmail(req.Email)
	if existEmail.ID != 0 {
		err := errors.New("email already exists")
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to register user")
		return err
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user.Password,_ = u.authUseCase.HashPassword(user.Password)
	return u.UserRepository.Create(user)
}

func (u *UserUsecaseImpl) Login(email, password string) (*string, error) {
	user, _ := u.UserRepository.FindByEmail(email)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	if !u.authUseCase.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	
	token, err := u.authUseCase.GenerateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &token, nil
}
