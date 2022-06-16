package serializer

import (
	"testing"

	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRegisteredResponse(t *testing.T) {
	asserts := assert.New(t)

	serializer := UserRegisteredSerializer{
		Ctx: nil,
		User: &model.User{
			Username: "miftah",
			UserRole: model.Role{Name: "User"},
		},
		Password: "123456",
	}
	response := serializer.Response()

	asserts.Equal(serializer.User.Username, response.Username)
	asserts.Equal(serializer.User.UserRole.Name, response.Role)
	asserts.Equal(serializer.Password, response.Password)
}
