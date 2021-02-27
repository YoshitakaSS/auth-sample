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

	return e
}
