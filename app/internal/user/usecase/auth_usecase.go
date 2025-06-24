package usecase

import "github.com/isaafisyah/order-management/app/internal/user/model"

type AuthUsecase interface {
	GenerateToken(user model.User) (string, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password string, hash string) bool
}