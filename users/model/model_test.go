package model

import (
	"fmt"
	"os"
	"testing"

	db "github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
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
	initTest()

	exit_code := m.Run()

	CleaningUpTest()
	db.CloseDatabase()
	os.Exit(exit_code)
	fmt.Println("Test User Model package end")
}

func TestSetPasswordHash(t *testing.T) {
	asserts := assert.New(t)

	mock_user := User{}
	asserts.NoError(mock_user.SetPasswordHash("password"), "should return no error")
	asserts.Error(mock_user.SetPasswordHash(""), "should return error empty password")
}

func TestCheckPassword(t *testing.T) {
	asserts := assert.New(t)

	mock_user := User{}
	mock_user.SetPasswordHash("password")
	asserts.NoError(mock_user.CheckPassWord("password"), "should return no error")
	asserts.Error(mock_user.CheckPassWord("123456"), "should return error")
}

func TestSaveOne(t *testing.T) {
	asserts := assert.New(t)

	// fmt.Println("user mock", usersMock[0])
	err := SaveOne(&UsersMock[0])
	asserts.NoError(err)
	// fmt.Println("user mock", usersMock[0])
	err = SaveOne(&User{Username: UsersMock[0].Username}) //create same user again
	asserts.Error(err, "Should return error duplicate key")
}

func TestFindOneUser(t *testing.T) {
	asserts := assert.New(t)

	//find existing user
	user, err := FindOneUser(&User{Username: UsersMock[0].Username})
	asserts.NoError(err)
	// fmt.Println("user", user)
	asserts.Equal(UsersMock[0].Username, user.Username, "Should return same user name")
	asserts.Equal(UsersMock[0].UserRole.Name, user.UserRole.Name, "Should return same user role")

	//find non existing user
	user, err = FindOneUser(&User{Username: "noName"})
	asserts.Error(err, "Should return error record not found")
	asserts.Empty(user.Username, "Should return empty user name")
	asserts.Empty(user.UserRole.Name, "Should return empty user role")
}
