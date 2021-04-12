package controller

import (
	"net/http"

	"github.com/YoshitakaSS/go_auth/presenter/request"
	"github.com/labstack/echo"
)

// UserRegister は User登録を行うContorller
func RegisterUser(c echo.Context) error {
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
