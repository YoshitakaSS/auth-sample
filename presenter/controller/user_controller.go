package controller

import (
	"net/http"

	"github.com/YoshitakaSS/go_auth/presenter/request"
	"github.com/labstack/echo"
)

// UserRegister は User登録を行うContorller
func UserRegister(c echo.Context) error {
	// requestBodyをbyteで持つ
	// b, err := ioutil.ReadAll(c.Request().Body)

	// if err != nil {
	// 	return err
	// }

	// if err := json.Unmarshal(b, r); err != nil {
	// 	panic(err)
	// }

	// new(request.UserRegister)
	// if err := json.Unmarshal(bodyBytes, r); err != nil {
	// 	return err
	// }
	r := &request.UserRegister{}

	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
