package main

import (
	"Mini_Project/configs"
	"Mini_Project/repository/auth"
	"Mini_Project/repository/user"

	// "Alterra/batch5/ORM/Layered/layered/delivery/controllers/auth"

	_authCon "Mini_Project/delivery/controllers/auth"
	_userCon "Mini_Project/delivery/controllers/user"
	"Mini_Project/delivery/routes"

	"Mini_Project/utils"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	userRepo := user.New(db)
	authRepo := auth.New(db)
	userController := _userCon.New(userRepo)
	authController := _authCon.New(authRepo)

	e := echo.New()

	routes.RegisterPath(e, userController, authController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
