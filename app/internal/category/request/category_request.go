package request

import "github.com/isaafisyah/order-management/app/internal/category/model"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"min=3,max=100"`
}


func ToCategory(category CreateCategoryRequest) model.Category {
	return model.Category{
		Name: category.Name,
	}
}

func ToCategoryUpdate(category UpdateCategoryRequest) model.Category {
	return model.Category{
		Name: category.Name,
	}
}