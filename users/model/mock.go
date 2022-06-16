package model

import (
	"fmt"

	db "github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/MiftahSalam/jabar-digital-service-test/users/dtos"
)

var UsersMock = []User{
	{
		Username: "miftah",
		Email:    "miftah@salam.com",
		UserRole: Role{
			Name: dtos.ROLES[0],
		},
	},
}

func CleaningUpTest() {
	fmt.Println("Clean up users start")

	db := db.Db.Connection

	//clean up users table
	for _, user := range UsersMock {
		db.Debug().Unscoped().Where("username = ?", user.Username).Delete(User{})
	}

	//clean up roles table
	for _, role := range dtos.ROLES {
		db.Unscoped().Where("name = ?", role).Delete(Role{})
	}

	fmt.Println("Clean up users end")
}

func initTest() {
	fmt.Println("Initialize users test start")

	for i := 0; i < len(UsersMock); i++ {
		UsersMock[i].SetPasswordHash("123456")
	}
	fmt.Println("Initialize users test end")
}
