package model

import (
	"errors"
	"fmt"

	"github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"github.com/MiftahSalam/jabar-digital-service-test/users/dtos"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null;unique"`
}

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;not null;unique"`
	Email        string `gorm:"column:email;not null"`
	PasswordHash string `gorm:"column:password;not null"`
	UserRole     Role   `gorm:"ForeignKey:RoleId"`
	RoleId       uint
}

func (user *User) TableName() string {
	return "users"
}

func (role Role) TableName() string {
	return "roles"
}

func Migrate() {
	database.Db.Connection.AutoMigrate(&Role{})
	database.Db.Connection.AutoMigrate(&User{})

	//init Role
	for _, role := range dtos.ROLES {
		var created_role Role
		err := database.Db.Connection.FirstOrCreate(&created_role, Role{Name: role}).Error
		if err != nil {
			fmt.Printf("Cannot create role %v with error %v\n", role, err)
		}
	}
}

func FindOneUser(condition interface{}, args ...interface{}) (User, error) {
	var user User
	err := database.Db.Connection.Preload("UserRole").Where(condition, args...).First(&user).Error

	return user, err
}

func CreateOneUser(data interface{}) error {
	if user_ref, ok := data.(*User); ok {
		user := *user_ref
		var role Role
		err := database.Db.Connection.First(&role, &Role{Name: user.UserRole.Name}).Error
		if err != nil {
			return err
		}
		user.UserRole = Role{}
		user.RoleId = role.Id

		return database.Db.Connection.Create(&user).Error
	}
	return database.Db.Connection.Save(data).Error
}

func (user *User) SetPasswordHash(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}

	byte_password := []byte(password)
	password_hash, _ := bcrypt.GenerateFromPassword(byte_password, bcrypt.DefaultCost)
	user.PasswordHash = string(password_hash)

	return nil
}

func (user *User) CheckPassWord(password string) error {
	byte_password := []byte(password)
	byte_password_hash := []byte(user.PasswordHash)

	return bcrypt.CompareHashAndPassword(byte_password_hash, byte_password)
}
