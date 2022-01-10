package user

import "Mini_Project/entities"

type User interface {
	GetAllUsersController() ([]entities.User, error)
}
