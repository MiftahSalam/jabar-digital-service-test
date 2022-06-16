package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/MiftahSalam/jabar-digital-service-test/commons"
	"github.com/MiftahSalam/jabar-digital-service-test/users/dtos"
	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	"github.com/MiftahSalam/jabar-digital-service-test/users/serializer"
)

func Register(ctx *gin.Context) {
	user := dtos.UserRegisterDto{}
	if err := user.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, commons.NewValidationError(err))
		return
	}

	found_user, _ := model.FindOneUser(&model.User{Username: user.Username})
	if found_user.Username != "" {
		ctx.JSON(http.StatusUnprocessableEntity, commons.NewError("database", errors.New("user already exist")))
		return
	} else {
		new_user := model.User{
			Username: user.Username,
			UserRole: model.Role{Name: user.Role},
		}
		password := commons.RandString(6)
		new_user.SetPasswordHash(password)
		model.SaveOne(&new_user)

		userSerializer := serializer.UserRegisteredSerializer{
			Ctx:      ctx,
			User:     &new_user,
			Password: password,
		}

		ctx.JSON(http.StatusCreated, gin.H{"user": userSerializer.Response()})
	}
}

func Login(ctx *gin.Context) {
	user := dtos.UserLoginDto{}
	if err := user.Bind(ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, commons.NewValidationError(err))
		return
	}

	found_user, _ := model.FindOneUser(&model.User{Username: user.Username})
	if found_user.Username == "" {
		ctx.JSON(http.StatusNotFound, commons.NewError("database", errors.New("user not found")))
		return
	} else {
		if found_user.CheckPassWord(user.Password) != nil {
			ctx.JSON(http.StatusForbidden, commons.NewError("login", errors.New("invalid password")))
			return
		}

		userSerializer := serializer.UserLoggedInSerializer{
			Ctx:  ctx,
			User: &found_user,
		}
		ctx.JSON(http.StatusOK, gin.H{"user": userSerializer.Response()})
	}
}

func Auth(ctx *gin.Context) {
	token_string := ctx.Param("token")

	// fmt.Println("token_string", token_string)

	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(token_string, claims, func(t *jwt.Token) (interface{}, error) {
		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusUnauthorized, commons.NewError("auth", err))
		return
	}

	token_claim := *claims
	// fmt.Println("token_claim", token_claim)

	auth_response := serializer.AuthDtoResponse{
		Is_valid: token.Valid,
		Expired:  time.UnixMilli(int64(token_claim["exp"].(float64))).String(),
		Username: token_claim["username"].(string),
	}

	ctx.JSON(http.StatusOK, gin.H{"validate": auth_response})
}
