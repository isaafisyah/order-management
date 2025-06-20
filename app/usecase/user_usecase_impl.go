package usecase

import (
	"errors"

	"github.com/isaafisyah/order-management/app/domain/model"
	"github.com/isaafisyah/order-management/app/domain/repository"
	"github.com/isaafisyah/order-management/app/domain/usecase"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &UserUsecaseImpl{userRepository}
}

func (u *UserUsecaseImpl) FindAll() ([]model.User, error) {
	return u.UserRepository.FindAll()
}

func (u *UserUsecaseImpl) FindByID(id int) (model.User, error) {
	return u.UserRepository.FindByID(id)
}

func (u *UserUsecaseImpl) FindByEmail(email string) (*model.User, error) {
	data, err := u.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserUsecaseImpl) Create(user model.User) error {
	existEmail, _ := u.UserRepository.FindByEmail(user.Email)
	if existEmail != nil {
		return errors.New("email already exists")
	}
	user.Password,_ = hashPassword(user.Password)
	return u.UserRepository.Create(user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
