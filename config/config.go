package config

import (
	"finalproject2/model"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadConfig() model.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbLoc := os.Getenv("DB_LOC")

	config := model.Config{
		ServerPort: serverPort,
		Database: model.Database{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Name:     dbName,
			Location: dbLoc,
		},
	}
	return config
}

func ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName, dbLoc string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", dbUsername, dbPassword, dbHost, dbPort, dbName, dbLoc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
