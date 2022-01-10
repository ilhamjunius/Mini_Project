package user

import (
	"Mini_Project/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAllUsersController() ([]entities.User, error) {
	var users []entities.User

	if err := ur.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
