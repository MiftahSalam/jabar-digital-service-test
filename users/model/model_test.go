package model

import (
	"fmt"
	"os"
	"testing"

	db "github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	fmt.Println("Test User Model package start")

	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Cannot load environment file")
	}

	err = db.InitDatabase()
	if err != nil {
		panic("Database error connection")
	}
	Migrate()

	exit_code := m.Run()

	os.Exit(exit_code)
	fmt.Println("Test User Model package end")

}
