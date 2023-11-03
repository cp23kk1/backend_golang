package databases

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	USERNAME string
	PASSWORD string
	NAME string
	HOST string
	PORT string
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	DB, _ = ConnectDB()
	return DB
}

func ConnectDB() (*gorm.DB, error) {
	config := LoadDatabaseConfig()
	var err error
	dsn := config.USERNAME +":"+ config.PASSWORD +"@tcp"+ "(" + config.HOST + ":" + config.PORT +")/" + config.NAME + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
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

// LoadDatabaseConfig loads the database configuration from environment variables.
func LoadDatabaseConfig() *Database {
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    return &Database{
        USERNAME: os.Getenv("DB_USERNAME"),
        PASSWORD: os.Getenv("DB_PASSWORD"),
        NAME: os.Getenv("DB_NAME"),
        HOST: os.Getenv("DB_HOST"),
        PORT: os.Getenv("DB_PORT"),
    }
}