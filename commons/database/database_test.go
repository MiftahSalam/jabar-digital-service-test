package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Test Database Commons package start")

	err := godotenv.Load("../../test.env")
	if err != nil {
		panic("Cannot load environment file")
	}

	exitCode := m.Run()

	CloseDatabase()
	os.Exit(exitCode)

	fmt.Println("Test Database Commons package end")
}
func TestInitDB(t *testing.T) {
	asserts := assert.New(t)

	asserts.Error(InitDatabase(), "Should return empty db port variable error")
	os.Setenv("DATABASE_PORT", "5432")
	asserts.Error(InitDatabase(), "Should return empty db host variable error")
	os.Setenv("DATABASE_HOST", "localhost")
	asserts.Error(InitDatabase(), "Should return empty db username variable error")
	os.Setenv("DATABASE_USERNAME", "postgres")
	asserts.Error(InitDatabase(), "Should return empty db name variable error")
	os.Setenv("DATABASE_NAME", "simple_api")
	asserts.NoError(InitDatabase(), "Should return nil error") //suppose simple_api db already exist
}
