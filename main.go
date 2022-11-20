package main

import (
	"fmt"
	"os"

	"github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/MiftahSalam/jabar-digital-service-test/users/model"
	userRouter "github.com/MiftahSalam/jabar-digital-service-test/users/router"
	"github.com/MiftahSalam/jabar-digital-service-test/users/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	production_mode := os.Getenv("PRODUCTION")
	if production_mode == "" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println("Error while loading env file with error", err)
			panic("Error while loading env file")
		}
	}

	err := database.InitDatabase()
	if err != nil {
		fmt.Println("Error while initialize database with error", err)
		panic("Error while initialize database")
	}

	validator.Init()
	model.Migrate() //migrate user models

	router := gin.Default()
	v1 := router.Group("/api/v1")
	userRouter.Users(v1.Group("/auth"))

	router.Run();
}
