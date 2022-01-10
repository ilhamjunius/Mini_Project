package user

import "Mini_Project/entities"

type CreateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GetUserResponseFormat struct {
	Data    []entities.User `json:"users"`
	Message string          `json:"message"`
}
type UserResponseFormat struct {
	Data    entities.User `json:"users"`
	Message string        `json:"message"`
}
