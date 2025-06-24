package request

import "github.com/isaafisyah/order-management/app/internal/product/model"

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	CategoryId  int    `json:"category_id" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

func ToProduct(product CreateProductRequest) model.Product {
	return model.Product{
		Name:        product.Name,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Stock: 		product.Stock,
	}
}