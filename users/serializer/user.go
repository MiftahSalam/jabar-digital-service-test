package serializer

import (
	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/gin-gonic/gin"
)

type UserRegisteredSerializer struct {
	Ctx      *gin.Context
	User     *model.User
	Password string
}

type UserRegisteredResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Password string `json:"password"`
}

func (user *UserRegisteredSerializer) Response() UserRegisteredResponse {
	response := UserRegisteredResponse{
		Username: user.User.Username,
		Password: user.Password,
		Role:     user.User.UserRole.Name,
	}

	return response
}
