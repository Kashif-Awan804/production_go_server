package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	ENV  string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT must be set")
	}

	ENV := os.Getenv("ENV")
	if ENV == "" {
		panic("ENV must be set")
	}

	return Config{
		PORT: PORT,
		ENV:  ENV,
	}

}
