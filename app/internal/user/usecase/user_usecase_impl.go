package usecase

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isaafisyah/order-management/app/config"
	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/isaafisyah/order-management/app/internal/user/repository"
	"github.com/isaafisyah/order-management/app/internal/user/response"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepository}
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.UserResponse{}, errors.New("user not found")
	}
	return response.ToUserResponse(data), nil
}

func (u *UserUsecaseImpl) Create(user model.User) error {
	existEmail, _ := u.FindByEmail(user.Email)
	if existEmail.ID != 0 {
		return errors.New("email already exists")
	}

	user.Password,_ = hashPassword(user.Password)
	return u.UserRepository.Create(user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (u *UserUsecaseImpl) Login(email, password string) (*string, error) {
	user, _ := u.UserRepository.FindByEmail(email)
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	
	token, err := generateToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}
	return &token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateToken(user model.User) (string, error) {
	cnf := config.Get()
	var secretKey = []byte(cnf.Server.SecretKey)

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
