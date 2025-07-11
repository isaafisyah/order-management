package repository

import (
	"context"

	"github.com/isaafisyah/order-management/app/internal/category/model"
	"gorm.io/gorm"
)


type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	return categories, err
}

func (r *CategoryRepositoryImpl) FindByID(ctx context.Context, id int) (model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).First(&category, id).Error
	return category, err
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, category model.Category) error {
	return r.db.WithContext(ctx).Create(&category).Error
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, category model.Category) error {
	return r.db.WithContext(ctx).Save(&category).Error
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, category model.Category) error {
	return r.db.WithContext(ctx).Delete(&category).Error
}