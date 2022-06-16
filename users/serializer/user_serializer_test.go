package serializer

import (
	"fmt"
	"testing"

	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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

func TestUserLoggedInResponse(t *testing.T) {
	asserts := assert.New(t)

	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Cannot load environment file")
	}

	serializer := UserLoggedInSerializer{
		Ctx: nil,
		User: &model.User{
			Username: "miftah",
			Model:    gorm.Model{ID: 20},
			UserRole: model.Role{Name: "User"},
		},
	}
	response := serializer.Response()

	fmt.Println("response", response)

	asserts.Equal(serializer.User.Username, response.Username)
	asserts.Equal(serializer.User.ID, response.ID)
	asserts.NotEmpty(response.Token)
}
