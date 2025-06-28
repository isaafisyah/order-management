package repository

import (
	"context"
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

func (r *UserRepositoryImpl) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Select([]string{"id", "name", "email"}).Find(&users).Error
	return users, err
	
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, id int) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return model.User{}, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepositoryImpl) Create(ctx context.Context,user model.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}
