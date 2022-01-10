package auth

import (
	"Mini_Project/entities"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a *AuthRepository) Login(email, password string) (entities.User, error) {
	foundUser := entities.User{}

	if err := a.db.Where("email = ? AND password = ?", email, password).Find(&foundUser).Error; err != nil {

		return foundUser, err
	}
	fmt.Println(foundUser)
	return foundUser, nil

}
