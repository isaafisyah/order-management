package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/delivery/http/handler"
	"github.com/isaafisyah/order-management/app/repository"
	"github.com/isaafisyah/order-management/app/usecase"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB,r *gin.Engine) {
	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(repository.NewUserRepository(db)))
	prefixGroup := r.Group("/api/v1")
	prefixGroup.GET("/users", userHandler.FindAll)
	prefixGroup.POST("/user", userHandler.Create)
	
}