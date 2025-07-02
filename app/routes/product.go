package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/internal/product/handler"
	"github.com/isaafisyah/order-management/app/internal/product/repository"
	"github.com/isaafisyah/order-management/app/internal/product/usecase"
	"gorm.io/gorm"
)

func ProductRoutes(db *gorm.DB,r *gin.Engine) {
	productHandler := handler.NewProductHandler(usecase.NewProductUsecase(repository.NewProductRepository(db)))
	prefixGroup := r.Group("/api/v1")
	prefixGroup.GET("/products", productHandler.FindAll)
	prefixGroup.POST("/product", productHandler.Create)
	prefixGroup.GET("/product/:id", productHandler.FindByID)
	prefixGroup.PUT("/product/:id", productHandler.Update)
	prefixGroup.DELETE("/product/:id", productHandler.Delete)
}