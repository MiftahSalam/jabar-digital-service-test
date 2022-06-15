package model

import (
	"errors"

	"github.com/MiftahSalam/jabar-digital-service-test/commons/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"column:name;not null"`
}

type User struct {
	gorm.Model
	Username     string `gorm:"column:username;not null;unique"`
	Email        string `gorm:"column:email;not null"`
	PasswordHash string `gorm:"column:password;not null"`
	UserRole     Role   `gorm:"ForeignKey:Id"`
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