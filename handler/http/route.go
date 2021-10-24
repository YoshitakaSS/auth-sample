package web

import (
	"net/http"

	"github.com/YoshitakaSS/go_auth/presenter/controller"
	"github.com/YoshitakaSS/go_auth/presenter/request"
	"github.com/labstack/echo"
)

// NewServer is return echo mapping
func NewServer() *echo.Echo {
	e := echo.New()
	e.Validator = request.NewValidator()

	// TOP
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TOP")
	})

	// API Group
	api := e.Group("/api")

	api.GET("/users/:user_name", func(c echo.Context) (err error) {
		return controller.FindUser(c)
	})

	// ユーザー登録
	api.POST("/users/regsiter", func(c echo.Context) (err error) {
		return controller.RegisterUser(c)
	})

	return e
}
