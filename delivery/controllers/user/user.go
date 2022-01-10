package user

import (
	"net/http"

	"Mini_Project/repository/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo user.User
}

func New(user user.User) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) GetAllUsersController() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.Repo.GetAllUsersController()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Something wrong",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users",
			"users":   res,
		})
	}

}
