package repository

import (
	"errors"

	"github.com/isaafisyah/order-management/app/internal/user/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select([]string{"id", "name", "email"}).Find(&users).Error
	return users, err
	
}

func (r *UserRepositoryImpl) FindByID(id int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepositoryImpl) Create(user model.User) error {
	return r.db.Create(&user).Error
}
