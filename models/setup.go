package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
