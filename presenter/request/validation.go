package request

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

// RequestValidator はValidatitorのポインタを持った構造体
type RequestValidator struct {
	validator *validator.Validate
}

// NewValidator RequestValidatorのポインタを返却
func NewValidator() echo.Validator {
	return &RequestValidator{validator: validator.New()}
}

// Validate validate
func (cv *RequestValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
