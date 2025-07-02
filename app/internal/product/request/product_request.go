package request

import "github.com/isaafisyah/order-management/app/internal/product/model"

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	CategoryId  int    `json:"category_id" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

type UpdateProductRequest struct{
	Name  string `json:"name" validate:"min=3,max=100"`
	CategoryId  int    `json:"category_id" validate:"required"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func ToProduct(product CreateProductRequest) model.Product {
	return model.Product{
		Name:        product.Name,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Stock: 		product.Stock,
	}
}

func UpdateToProduct(id int ,product UpdateProductRequest) model.Product {
	return model.Product{
		ID:          id,
		Name:        product.Name,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Stock: 		product.Stock,
	}
}