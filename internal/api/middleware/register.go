package middleware

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterJSONTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		engine, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			engine.RegisterTagNameFunc(func(field reflect.StructField) string {
				name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
				if name == "-" {
					return ""
				}

				return name
			})
		}

		c.Next()
	}
}
