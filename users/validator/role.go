package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/MiftahSalam/jabar-digital-service-test/users/dtos"
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", validateRole)
	}
}

func validateRole(field validator.FieldLevel) bool {
	role := field.Field().String()
	for _, v := range dtos.ROLES {
		if role == v {
			return true
		}
	}
	return false
}
