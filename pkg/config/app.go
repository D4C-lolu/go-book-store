package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("C:\\Users\\Demilade\\Documents\\web_projects\\Golang\\book_store\\.env")
	if err != nil {
		fmt.Printf("%s", err)
		panic(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	dbString := fmt.Sprintf("%s:%s@tcp(%s)", dbUser, dbPassword, dbURL)
	d, err := gorm.Open("mysql", fmt.Sprintf("%s/%s?charset=utf8&parseTime=True&loc=Local", dbString, dbName))
	if err != nil {
		fmt.Printf("%s", err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
