package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT          string
	ServerAddress string
	WhiteListed   string
}

var config Config

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("ERROR .env Not found")
	}

	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	config.PORT = os.Getenv("PORT")
	config.WhiteListed = os.Getenv("WHITELISTED_URLS")
}

func GetConfig() *Config {
	return &config
}
