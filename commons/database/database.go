package database

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB //shared database connection global variable
	Host       string
	Username   string
	Name       string
	Password   string
	Url        string
	Port       int
}

var Db Database

//Initiate database connection with parameter from environment variabel
//function will return related error
func InitDatabase() error {
	production_mode := os.Getenv("PRODUCTION")

	if production_mode == "" {
		db_port_str := os.Getenv("DATABASE_PORT")
		if db_port_str == "" {
			fmt.Println("Error while loading environment variable DATABASE_PORT")
			return errors.New("empty db port variable")
		}

		db_port, err := strconv.Atoi(db_port_str)
		if err != nil {
			fmt.Println("Error while loading environment variable DATABASE_PORT")
			return err
		} else {
			Db.Port = db_port
		}

		db_host := os.Getenv("DATABASE_HOST")
		if db_host == "" {
			fmt.Println("Error while loading environment variable DATABASE_HOST")
			return errors.New("empty db host variable")
		} else {
			Db.Host = db_host
		}

		db_username := os.Getenv("DATABASE_USERNAME")
		if db_username == "" {
			fmt.Println("Error while loading environment variable DATABASE_USERNAME")
			return errors.New("empty db username variable")
		} else {
			Db.Username = db_username
		}

		db_password := os.Getenv("DATABASE_PASSWORD")
		if db_password == "" {
			fmt.Println("Error while loading environment variable DATABASE_PASSWORD")
			return errors.New("empty db password variable")
		} else {
			Db.Password = db_password
		}

		db_name := os.Getenv("DATABASE_NAME")
		if db_name == "" {
			fmt.Println("Error while loading environment variable DATABASE_NAME")
			return errors.New("empty db name variable")
		} else {
			Db.Name = db_name
		}

		Db.Url = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
			Db.Host,
			Db.Username,
			Db.Password,
			Db.Name,
			Db.Port)
	} else {
		db_url := os.Getenv("DATABASE_URL")
		if db_url == "" {
			fmt.Println("Error while loading environment variable DATABASE_URL")
			return errors.New("empty db url variable")
		} else {
			Db.Url = db_url //heroku environment variable
		}
	}

	db, err := gorm.Open(postgres.Open(Db.Url))
	if err != nil {
		fmt.Println("Database Error", err)
		return err
	} else {
		Db.Connection = db
	}

	sqlDb, err := Db.Connection.DB()
	if err != nil {
		fmt.Println("Database Error", err)
		return err
	}
	sqlDb.SetConnMaxIdleTime(10)

	return nil
}

func CloseDatabase() {
	if Db.Connection != nil {
		sqlDB, err := Db.Connection.DB()
		if err != nil {
			fmt.Println("Cannot get database instance")
		} else {
			sqlDB.Close()
		}
	} else {
		fmt.Println("Cannot get database connection")
	}
}
