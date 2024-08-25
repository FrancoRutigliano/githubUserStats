package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() (string, error) {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		return "", errors.New("error to load the port variable")
	}

	return port, nil
}
