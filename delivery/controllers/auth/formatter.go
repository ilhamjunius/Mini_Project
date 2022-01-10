package auth

import "Mini_Project/entities"

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginResponseFormat struct {
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
	Token   string        `json:"token"`
}
