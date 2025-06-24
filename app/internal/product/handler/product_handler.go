package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/internal/product/request"
	"github.com/isaafisyah/order-management/app/internal/product/usecase"
	"github.com/isaafisyah/order-management/app/utils/logger"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) ProductHandler {
	return ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (h *ProductHandler) FindAll(ctx *gin.Context) {
	logger.Log.Info("Find all products")
	users, err := h.ProductUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "ProductHandler").WithError(err).Error("Failed to find all products")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (h *ProductHandler) FindByID(ctx *gin.Context) {
	logger.Log.Info("Find product by id")
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	user, err := h.ProductUsecase.FindByID(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "ProductHandler").WithError(err).Error("Failed to find product by id")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	logger.Log.Info("Create product")
	var req request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "ProductHandler").WithError(err).Error("Failed to create product")
		return
	}

	if err := h.ProductUsecase.Create(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "ProductHandler").WithError(err).Error("Failed to create product")
		return
	}
	logger.Log.Info("Product created successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
	})
}