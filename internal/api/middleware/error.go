package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/elgntt/notes/internal/api/validation"
)

func ValidationErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, e := range c.Errors {
			switch e.Type {
			case gin.ErrorTypeBind:
				var jErr *json.UnmarshalTypeError
				if errors.As(e.Err, &jErr) {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"error": validation.FieldError{
							Field: jErr.Field,
							Error: fmt.Sprintf(
								"invalid data type for field '%s': expected %s",
								jErr.Field,
								jErr.Type.String(),
							),
						},
					})
					return
				}

				var vErr validator.ValidationErrors
				if errors.As(e.Err, &vErr) {
					errs := make([]validation.FieldError, 0, len(vErr))
					for _, fe := range vErr {
						errs = append(errs, validation.FieldError{
							Field: fe.Field(),
							Error: validation.FieldErrorToText(fe),
						})
					}

					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"errors": errs,
					})
					return
				}

				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": e.Error(),
				})
				return
			}
		}

	}
}
