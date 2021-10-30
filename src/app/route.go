package web

import (
	"net/http"

	"github.com/labstack/echo"
)

// NewServer is return echo mapping
func NewServer() *echo.Echo {
	e := echo.New()

	// TOP
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "TOP")
	})

	// API Group
	api := e.Group("/api")
	// ユーザー登録
	api.POST("/users/regsiter", func(c echo.Context) (err error) {
		return nil
	})

	return e
}
