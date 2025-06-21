package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/isaafisyah/order-management/app/internal/user/model"
	"github.com/isaafisyah/order-management/app/internal/user/request"
	"github.com/isaafisyah/order-management/app/internal/user/usecase"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/isaafisyah/order-management/app/utils/validator"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return UserHandler{
		UserUsecase: userUsecase,
	}
}

func (c *UserHandler) FindAll(ctx *gin.Context) {
	logger.Log.Info("Find all users")
	users, err := c.UserUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to find all users")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (c *UserHandler) FindByID(ctx *gin.Context) {
	logger.Log.Info("Find user by id")
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	user, err := c.UserUsecase.FindByID(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to find user by id")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c *UserHandler) Register(ctx *gin.Context) {
	logger.Log.Info("Register user")
	var req request.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to register user")
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to register user")
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := c.UserUsecase.Create(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to register user")
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (c *UserHandler) Login(ctx *gin.Context)  {
	logger.Log.Info("Login user")
	var req request.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to login user")
		return
	}

	if err := validator.ValidateStruct(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to login user")
		return
	}

	token, err := c.UserUsecase.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		logger.Log.WithField("Module", "UserHandler").WithError(err).Error("Failed to login user")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}