package auth

import (
	"Mini_Project/entities"
)

type Auth interface {
	Login(name, password string) (entities.User, error)
}
