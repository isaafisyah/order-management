package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/internal/user/handler"
	"github.com/isaafisyah/order-management/app/internal/user/repository"
	"github.com/isaafisyah/order-management/app/internal/user/usecase"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB,r *gin.Engine) {
	userHandler := handler.NewUserHandler(usecase.NewUserUsecase(repository.NewUserRepository(db)))
	prefixGroup := r.Group("/api/v1")
	prefixGroup.GET("/users", userHandler.FindAll)
	prefixGroup.POST("/register", userHandler.Register)
	prefixGroup.POST("/login", userHandler.Login)
	prefixGroup.GET("/user/:id", userHandler.FindByID)
	
}