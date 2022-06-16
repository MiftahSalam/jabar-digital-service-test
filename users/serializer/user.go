package serializer

import (
	"github.com/MiftahSalam/jabar-digital-service-test/commons"
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

type UserLoggedInSerializer struct {
	Ctx  *gin.Context
	User *model.User
}

type UserLoggedInResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (user *UserRegisteredSerializer) Response() UserRegisteredResponse {
	response := UserRegisteredResponse{
		Username: user.User.Username,
		Password: user.Password,
		Role:     user.User.UserRole.Name,
	}

	return response
}

func (user *UserLoggedInSerializer) Response() UserLoggedInResponse {
	response := UserLoggedInResponse{
		Username: user.User.Username,
		ID:       user.User.ID,
		Token:    commons.GenerateToken(user.User.ID, user.User.Username),
	}

	return response
}
