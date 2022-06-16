package service

import (
	"errors"
	"net/http"

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
