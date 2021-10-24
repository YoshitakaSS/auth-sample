package controller

import (
	"net/http"

	"github.com/YoshitakaSS/go_auth/presenter/request"
	"github.com/YoshitakaSS/go_auth/usecase"
	"github.com/labstack/echo"
)

type UserContorller struct {
	ufuc *usecase.UserFindUseCase
}

// UserRegister は User登録を行うContorller
func (c *UserContorller) RegisterUser(c echo.Context) error {
	r := &request.UserRegisterRequest{}

	// validate := validator.New()
	// if err := validate.Struct(r); err != nil {
	// 	validationErrors := err.(validator.ValidationErrors)
	// 	echo.NewHTTPError(http.StatusBadRequest, validationErrors)
	// }

	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}

func (c *userController) FindUser(c echo.Context) error {
	user_name := c.ParamValues()

	return c.JSON(http.StatusOK, user_name)
}
