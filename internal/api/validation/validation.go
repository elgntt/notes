package validation

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/elgntt/notes/internal/model/dto"
)

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateTaskStatus(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(dto.Enum).IsValid()
}

func ValidateDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().Interface().(dto.Date)

	_, err := time.Parse(dto.CustomDate, strings.Trim(dateStr.Format(dto.CustomDate), "\""))
	if err != nil {
		return false
	}

	return true
}

func Set() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("status", ValidateTaskStatus)
		if err != nil {
			return err
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("date", ValidateDate)
		if err != nil {
			return err
		}
	}

	return nil
}

func FieldErrorToText(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return fmt.Sprintf("This field cannot be longer than %s", e.Param())
	case "min":
		return fmt.Sprintf("This field  must be longer than %s", e.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("This field must be %s characters long", e.Param())
	case "status":
		return "This field should have one of these values: backlog, open, progress, review, completed."
	case "date":
		return "Incorrect date format, expected format: day-month-year (for example, 01-02-2006)"
	}

	return "This field is not valid"
}
