package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func LoadConfig() DBConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	return DBConfig{
		Host: 		os.Getenv("DB_HOST"),
		Port: 		os.Getenv("DB_PORT"),
		Username: 	os.Getenv("DB_USERNAME"),
		Password: 	os.Getenv("DB_PASSWORD"),
		DBName: 	os.Getenv("DB_NAME"),
	}
}
