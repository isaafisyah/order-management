package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/order-management/app/internal/category/request"
	"github.com/isaafisyah/order-management/app/internal/category/usecase"
	"github.com/isaafisyah/order-management/app/utils/logger"
)


type CategoryHandler struct {
	CategoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) CategoryHandler {
	return CategoryHandler{
		CategoryUsecase: categoryUsecase,
	}
}

func (r CategoryHandler) FindAll(ctx *gin.Context) {
	logger.Log.Info("Find all categories")
	categories, err := r.CategoryUsecase.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to find all categories")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}

func (r CategoryHandler) FindByID(ctx *gin.Context) {
	logger.Log.Info("Find category by id")
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	category, err := r.CategoryUsecase.FindByID(ctx.Request.Context(), idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to find category by id")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}

func (r CategoryHandler) Create(ctx *gin.Context) {
	logger.Log.Info("Create category")
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to create category")
		return
	}

	if err := r.CategoryUsecase.Create(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to create category")
		return
	}
	logger.Log.Info("Category created successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
	})
}

func (r CategoryHandler) Update(ctx *gin.Context){
	logger.Log.Info("Update Category")
	var req request.UpdateCategoryRequest
	if err := ctx.ShouldBindBodyWithJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to Update Category")
	}

	if err := r.CategoryUsecase.Update(ctx.Request.Context(), req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to update category")
		return
	}
	logger.Log.Info("Update Category successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category update successfully",
	})
}

func (r CategoryHandler) Delete(ctx *gin.Context) {
	logger.Log.Info("Delete Category")
	id,_ := strconv.Atoi(ctx.Param("id"))
	if err := r.CategoryUsecase.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "CategoryHandler").WithError(err).Error("Failed to delete category")
		return
	}
	logger.Log.Info("Category deleted successfully")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}