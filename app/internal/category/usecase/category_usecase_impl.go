package usecase

import (
	"context"
	"errors"

	"github.com/isaafisyah/order-management/app/internal/category/model"
	"github.com/isaafisyah/order-management/app/internal/category/repository"
	"github.com/isaafisyah/order-management/app/internal/category/request"
	"github.com/isaafisyah/order-management/app/utils/logger"
	"github.com/isaafisyah/order-management/app/utils/validator"
)

type CategoryUsecaseImpl struct{
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{
		repo: repo,
	}
}

func (r *CategoryUsecaseImpl) FindAll(ctx context.Context) ([]model.Category, error) {
	categories, err := r.repo.FindAll(ctx)
	return categories, err

}

func (r *CategoryUsecaseImpl) FindByID(ctx context.Context, id int) (model.Category, error) {
	return r.repo.FindByID(ctx, id)
}

func (r *CategoryUsecaseImpl) Create(ctx context.Context, req request.CreateCategoryRequest) error {
	if err := validator.ValidateStruct(&req); err != nil {
		err = errors.New("validation error")
		logger.Log.WithField("Module", "Category ProductService").WithError(err).Error("Failed to update category")
		return err
	}
	category := request.ToCategory(req)
	return r.repo.Create(ctx, category)
}

func (r *CategoryUsecaseImpl) Update(ctx context.Context, req request.UpdateCategoryRequest) error {
	if err := validator.ValidateStruct(&req); err != nil {
		err = errors.New("validation error")
		logger.Log.WithField("Module", "Category Service").WithError(err).Error("Failed to update category")
		return err
	}
	category := request.ToCategoryUpdate(req)
	return r.repo.Update(ctx, category)
}

func (r *CategoryUsecaseImpl) Delete(ctx context.Context, id int) error {
	category,_ := r.repo.FindByID(ctx, id)
	return r.repo.Delete(ctx, category)
}