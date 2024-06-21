package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	validatorpkg "task-manager/internal/pkg/validator"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  any    `json:"error"`
}

type ValidationErrors struct {
	Status string       `json:"status"`
	Errors []FieldError `json:"errors"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func WithJSONError(c *gin.Context, err error) {
	var jSyntaxErr *json.SyntaxError
	if errors.As(err, &jSyntaxErr) {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status: "error",
			Error: Error{
				Message: jSyntaxErr.Error(),
			},
		})

		return
	}

	var jErr *json.UnmarshalTypeError
	if errors.As(err, &jErr) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": FieldError{
				Field: jErr.Field,
				Message: fmt.Sprintf(
					"invalid data type for field '%s': expected %s",
					jErr.Field,
					jErr.Type.String(),
				),
			},
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	return
}

func WithValidationError(c *gin.Context, err error) {
	var vErr validator.ValidationErrors
	if errors.As(err, &vErr) {
		errs := make([]FieldError, 0, len(vErr))
		for _, fe := range vErr {
			errs = append(errs, FieldError{
				Field:   fe.Field(),
				Message: validatorpkg.FieldErrorToText(fe),
			})
		}

		c.JSON(http.StatusUnprocessableEntity, ValidationErrors{
			Status: "errors",
			Errors: errs,
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})

	return
}

func WithInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Something went wrong, try again later.",
	})
}

func WithNotFoundErr(c *gin.Context, errMsg string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": errMsg,
	})
}

func WithParameterInvalidErr(c *gin.Context, errMsg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": errMsg,
	})
}
