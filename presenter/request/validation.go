package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

// RequestValidator
type RequestValidator struct {
	validator *validator.Validate
}

// NewValidator
func NewValidator() echo.Validator {
	return &RequestValidator{validator: validator.New()}
}

// Validate validate
func (cv *RequestValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
