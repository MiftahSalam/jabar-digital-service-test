package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var ROLES = []string{"User", "Admin"}

type UserRegisterDto struct {
	Username string `json:"username" binding:"required,min=4,alphanum,max=50"`
	Role     string `json:"role" binding:"required,role"`
}

type UserLoginDto struct {
	Username string `json:"username" binding:"required,min=4,alphanum,max=50"`
	Password string `json:"password" binding:"required"`
}

func (user *UserRegisterDto) Bind(ctx *gin.Context) error {
	bind := binding.Default(ctx.Request.Method, ctx.ContentType())
	err := ctx.ShouldBindWith(user, bind)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (user *UserLoginDto) Bind(ctx *gin.Context) error {
	bind := binding.Default(ctx.Request.Method, ctx.ContentType())
	err := ctx.ShouldBindWith(user, bind)
	if err != nil {
		return err
	} else {
		return nil
	}
}
