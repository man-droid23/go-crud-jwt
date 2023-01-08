package response

import "fiber-rest/models/entity"

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

func NewUserListResponse(users []entity.User) []UserResponse {
	var userList []UserResponse
	for _, user := range users {
		userList = append(userList, NewUserResponse(user))
	}
	return userList
}
