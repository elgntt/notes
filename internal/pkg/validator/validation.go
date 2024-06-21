package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"task-manager/internal/model/dto"
)

type CustomValidator struct {
	*validator.Validate
}

func New() (*CustomValidator, error) {
	cv := &CustomValidator{
		validator.New(),
	}

	cv.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	if err := cv.RegisterValidation("status", isValidStatus); err != nil {
		return nil, fmt.Errorf("failed to register validator: %w", err)
	}

	return cv, nil
}

func isValidStatus(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(dto.Enum).IsValid()
}

func FieldErrorToText(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("field %s must be at least %s characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("field %s must be at most %s characters", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("field %s must be greater than or equal to %s", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("field %s must be less than or equal to %s", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("Invalid email format")
	case "len":
		return fmt.Sprintf("This field must be %s characters long", fe.Param())
	case "status":
		return "This field should have one of these values: backlog, open, progress, review, completed."
	}

	return "This field is not valid"
}
