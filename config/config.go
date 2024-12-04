package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Database Database
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
		Schema   string
	}
)

func LoadConfig() Config {

	if ex := godotenv.Load(".env"); ex != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Database: Database{
			Host:     os.Getenv(`PG_HOST`),
			Port:     os.Getenv(`PG_PORT`),
			User:     os.Getenv(`PG_USER`),
			Password: os.Getenv(`PG_PASSWORD`),
			DBName:   os.Getenv(`PG_DB_NAME`),
			SSLMode:  os.Getenv(`PG_SSL_MODE`),
			Schema:   os.Getenv(`PG_SCHEMA`),
		},
	}
}
