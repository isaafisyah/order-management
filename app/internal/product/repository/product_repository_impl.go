package repository

import (
	"github.com/isaafisyah/order-management/app/internal/product/model"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) FindByID(id int) (model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return product, err
}

func (r *ProductRepositoryImpl) Create(product model.Product) error {
	return r.db.Create(&product).Error
}

func (r *ProductRepositoryImpl) Update(product model.Product) error {
	return r.db.Save(&product).Error
}

func (r *ProductRepositoryImpl) Delete(product model.Product) error {
	return r.db.Delete(&product).Error
	
}

