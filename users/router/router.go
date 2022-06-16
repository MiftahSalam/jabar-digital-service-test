package router

import (
	"github.com/MiftahSalam/jabar-digital-service-test/users/service"
	"github.com/gin-gonic/gin"
)

func Users(router *gin.RouterGroup) {
	router.POST("/", service.Register)
	router.POST("/login", service.Login)
	router.GET("/:token", service.Auth)
}
