package my

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// CustomValidator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	fmt.Println("--- Validate Call")
	fmt.Println(i)

	cv.validator.RegisterValidation("print", MyValidate)

	return cv.validator.Struct(i)
}

func MyValidate(fl validator.FieldLevel) bool {
	fmt.Println("--- MyValidate Call")
	fmt.Printf("Field %T: %v\n", fl.Field(), fl.Field().String())
	return true
}