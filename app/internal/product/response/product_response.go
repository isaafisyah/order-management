package response

import "github.com/isaafisyah/order-management/app/internal/product/model"

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CategoryId  int    `json:"category_id"`
	Price       int    `json:"price"`
	Stock 		int    `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToProductResponse(product model.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		CategoryId:  product.CategoryId,
		Price:       product.Price,
		Stock: 		product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func ToProductResponses(product []model.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range product {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
	
}