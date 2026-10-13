package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	Port string
	DNS  string
}

var Env *env

func LoadConfig() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	Env = &env{
		Port: os.Getenv("APP_PORT"),
		DNS:  os.Getenv("DB_DSN"),
	}
}
