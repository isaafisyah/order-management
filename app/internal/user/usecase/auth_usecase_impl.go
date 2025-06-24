package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isaafisyah/order-management/app/config"
	"github.com/isaafisyah/order-management/app/internal/user/model"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {}

func NewAuthUsecase() AuthUsecase {
	return &AuthUsecaseImpl{}
}


func (a *AuthUsecaseImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (a *AuthUsecaseImpl) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *AuthUsecaseImpl) GenerateToken(user model.User) (string, error) {
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




