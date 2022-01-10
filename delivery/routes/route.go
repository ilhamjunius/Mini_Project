package routes

import (
	"Mini_Project/delivery/controllers/auth"
	"Mini_Project/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uc *user.UserController, ac *auth.AuthController) {

	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/users", uc.GetAllUsersController())

	e.Logger.Fatal(e.Start(":8000"))

}
