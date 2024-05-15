package databases

import (
	"cp23kk1/common/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	USERNAME string
	PASSWORD string
	NAME     string
	HOST     string
	PORT     string
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	DB, _ = ConnectDB()
	return DB
}

func ConnectDB() (*gorm.DB, error) {
	config, err := config.LoadConfig()
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp" + "(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=", err)
		return nil, err
	}
	fmt.Println("Succesfully connect the database!")
	return db, nil
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
