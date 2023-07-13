package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret		string
	JWTExpiration	time.Duration
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	
	return Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpiration: time.Hour * 24,
	}
 
}
