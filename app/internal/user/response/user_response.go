package response

import "github.com/isaafisyah/order-management/app/internal/user/model"

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserResponses(users []model.User) []UserResponse {
	var userResponse []UserResponse
	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}
	return userResponse
}