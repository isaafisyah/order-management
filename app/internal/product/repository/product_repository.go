package repository

import "github.com/isaafisyah/order-management/app/internal/product/model"

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id int) (model.Product, error)
	Create(product model.Product) error
	Update(product model.Product) error
}